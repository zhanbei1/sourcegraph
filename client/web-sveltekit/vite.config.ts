import { join } from 'path'

import { sveltekit } from '@sveltejs/kit/vite'
import { defineConfig, mergeConfig, type UserConfig } from 'vite'
import inspect from 'vite-plugin-inspect'

import graphqlCodegen from './dev/vite-graphql-codegen'

export default defineConfig(({ mode }) => {
    let config: UserConfig = {
        plugins: [
            sveltekit(),
            // Generates typescript types for gql-tags and .gql files
            graphqlCodegen(),
            inspect(),
        ],
        define:
            mode === 'test'
                ? {}
                : {
                      'process.platform': '"browser"',
                      'process.env.VITEST': 'undefined',
                      'process.env': '({})',
                  },
        css: {
            preprocessorOptions: {
                scss: {
                    includePaths: [
                        // Our scss files and scss files in client/* often import global styles via @import 'wildcard/src/...'
                        // Adding '..' as load path causes scss to look for these imports in the client folder.
                        // (without it scss @import paths are always relative to the importing file)
                        join(__dirname, '..'),
                    ],
                },
            },
            modules: {
                localsConvention: 'camelCase',
            },
        },
        server: {
            // Allow setting the port via env variables to make it easier to integrate with
            // our existing caddy setup (which proxies requests to a specific port).
            port: process.env.SK_PORT ? +process.env.SK_PORT : undefined,
            strictPort: !!process.env.SV_PORT,
            proxy: {
                // Proxy requests to specific endpoints to a real Sourcegraph
                // instance.
                '^(/sign-in|/.assets|/-|/.api|/search/stream|/users|/notebooks|/insights)': {
                    target: process.env.SOURCEGRAPH_API_URL || 'https://sourcegraph.com',
                    changeOrigin: true,
                    secure: false,
                },
            },
        },

        resolve: {
            alias: [
                // Unclear why Vite fails. It claims that index.esm.js doesn't have this export (it does).
                // Rewriting this to index.js fixes the issue. Error:
                // import { CiWarning, CiSettings, CiTextAlignLeft } from "react-icons/ci/index.esm.js";
                //                     ^^^^^^^^^^
                // SyntaxError: Named export 'CiSettings' not found. The requested module 'react-icons/ci/index.esm.js'
                // is a CommonJS module, which may not support all module.exports as named exports.
                {
                    find: /^react-icons\/(.+)$/,
                    replacement: 'react-icons/$1/index.js',
                },
                // We generate corresponding .gql.ts files for .gql files.
                // This alias allows us to import .gql files and have them resolved to the generated .gql.ts files.
                {
                    find: /^(.*)\.gql$/,
                    replacement: '$1.gql.ts',
                },
            ],
        },

        optimizeDeps: {
            exclude: [
                // Without addings this Vite throws an error
                'linguist-languages',
            ],
        },

        test: {
            setupFiles: './src/testing/setup.ts',
            include: ['src/**/*.test.ts'],
        },
    }

    if (process.env.BAZEL) {
        // Merge settings necessary to make the build work with bazel
        config = mergeConfig(config, {
            resolve: {
                alias: [
                    // When using Bazel, @sourcegraph/* dependencies will refer to the built packages.
                    // These do not contain the source *.module.scss files but still contain import statements
                    // that reference *.scss files. Processing them with vite throws an error unless we
                    // update the imports to reference the corresponding *.css files instead.
                    // Additionally our own source files might reference *.module.scss files, which we also want
                    // to rewrite.
                    {
                        find: /^(.+)\.module\.scss$/,
                        replacement: '$1.module.css',
                        customResolver(source, importer, options) {
                            // The this.resolve(...) part is taken from the @rollup/plugin-alias implementation. Without
                            // it it appears the bundler tries to resolve relative module IDs to the current working
                            // directory.
                            return source.includes('@sourcegraph') || importer?.includes('@sourcegraph/')
                                ? this.resolve(source, importer, { skipSelf: true, ...options }).then(
                                      resolved => resolved || { id: source }
                                  )
                                : null
                        },
                    },
                ],
            },
            ssr: {
                // By default vite treats dependencies that are links to other packages in the monorepo as source code
                // and processes them as well.
                // In a bazel sandbox however all @sourcegraph/* dependencies are built packages and thus not processed
                // by vite without this additional setting.
                // We have to process those files to apply certain "fixes", such as aliases defined in svelte.config.js.
                noExternal: [/@sourcegraph\/.*/],
            },
        } satisfies UserConfig)
    }

    return config
})
