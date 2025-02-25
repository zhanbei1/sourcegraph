import { FC, useMemo } from 'react'

import { FilterType, resolveFilter } from '@sourcegraph/shared/src/search/query/filters'
import { findFilters } from '@sourcegraph/shared/src/search/query/query'
import { scanSearchQuery, succeedScan } from '@sourcegraph/shared/src/search/query/scanner'
import type { Filter as QueryFilter } from '@sourcegraph/shared/src/search/query/token'
import { omitFilter, updateFilter } from '@sourcegraph/shared/src/search/query/transformer'
import type { Filter } from '@sourcegraph/shared/src/search/stream'

import {
    authorFilter,
    commitDateFilter,
    languageFilter,
    repoFilter,
    SearchDynamicFilter,
    symbolFilter,
    utilityFilter,
} from './components/dynamic-filter/SearchDynamicFilter'
import {
    FilterTypeList,
    resolveFilterTypeValue,
    toSearchSyntaxTypeFilter,
} from './components/filter-type-list/FilterTypeList'
import { FiltersDocFooter } from './components/filters-doc-footer/FiltersDocFooter'
import { useUrlFilters } from './hooks'
import { SearchFilterType } from './types'

import styles from './NewSearchFilters.module.scss'

interface NewSearchFiltersProps {
    query: string
    filters?: Filter[]
    onQueryChange: (nextQuery: string) => void
}

export const NewSearchFilters: FC<NewSearchFiltersProps> = ({ query, filters, onQueryChange }) => {
    const [selectedFilters, setSelectedFilters] = useUrlFilters()

    const type = useMemo(() => {
        const tokens = scanSearchQuery(query)

        if (tokens.type === 'success') {
            const filters = tokens.term.filter((token): token is QueryFilter => token.type === 'filter')
            const typeFilters = filters.filter(filter => resolveFilter(filter.field.value)?.type === 'type')

            if (typeFilters.length === 0 || typeFilters.length > 1) {
                return SearchFilterType.Code
            }

            return resolveFilterTypeValue(typeFilters[0].value?.value)
        }

        return SearchFilterType.Code
    }, [query])

    const handleFilterTypeChange = (filterType: SearchFilterType): void => {
        switch (filterType) {
            case SearchFilterType.Code: {
                const filters = findFilters(succeedScan(query), FilterType.type)

                const newQuery = filters.reduce((query, filter) => omitFilter(query, filter), query)
                onQueryChange(newQuery)
                break
            }
            default: {
                const filters = findFilters(succeedScan(query), FilterType.type)
                const newQuery = filters.reduce((query, filter) => omitFilter(query, filter), query)

                onQueryChange(updateFilter(newQuery, FilterType.type, toSearchSyntaxTypeFilter(filterType)))
            }
        }
    }

    return (
        <aside className={styles.scrollWrapper}>
            <FilterTypeList value={type} onSelect={handleFilterTypeChange} />

            {type === SearchFilterType.Symbols && (
                <SearchDynamicFilter
                    title="By symbol kind"
                    filterKind="symbol type"
                    filters={filters}
                    selectedFilters={selectedFilters}
                    renderItem={symbolFilter}
                    onSelectedFilterChange={setSelectedFilters}
                />
            )}

            {(type === SearchFilterType.Commits || type === SearchFilterType.Diffs) && (
                <>
                    <SearchDynamicFilter
                        title="By author"
                        filterKind="author"
                        filters={filters}
                        selectedFilters={selectedFilters}
                        renderItem={authorFilter}
                        onSelectedFilterChange={setSelectedFilters}
                    />

                    <SearchDynamicFilter
                        title="By commit date"
                        filterKind="commit date"
                        filters={filters}
                        selectedFilters={selectedFilters}
                        renderItem={commitDateFilter}
                        onSelectedFilterChange={setSelectedFilters}
                    />
                </>
            )}

            <SearchDynamicFilter
                title="By language"
                filterKind="lang"
                filters={filters}
                selectedFilters={selectedFilters}
                renderItem={languageFilter}
                onSelectedFilterChange={setSelectedFilters}
            />

            <SearchDynamicFilter
                title="By repositories"
                filterKind="repo"
                filters={filters}
                selectedFilters={selectedFilters}
                renderItem={repoFilter}
                onSelectedFilterChange={setSelectedFilters}
            />

            <SearchDynamicFilter
                title="By file"
                filterKind="file"
                filters={filters}
                selectedFilters={selectedFilters}
                onSelectedFilterChange={setSelectedFilters}
            />

            <SearchDynamicFilter
                title="Utility"
                filterKind="utility"
                filters={filters}
                selectedFilters={selectedFilters}
                renderItem={utilityFilter}
                onSelectedFilterChange={setSelectedFilters}
            />

            <FiltersDocFooter className={styles.footer} />
        </aside>
    )
}
