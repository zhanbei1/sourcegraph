import { ReactElement, useEffect } from 'react'

import { TelemetryProps } from '@sourcegraph/shared/src/telemetry/telemetryService'
import { Text } from '@sourcegraph/wildcard'

import { SiteAdminRepositoriesContainer } from '../../site-admin/SiteAdminRepositoriesContainer'

import { CustomNextButton } from './setup-steps'

interface SyncRepositoriesStep extends TelemetryProps {}

export function SyncRepositoriesStep(props: SyncRepositoriesStep): ReactElement {
    const { telemetryService } = props

    useEffect(() => {
        telemetryService.log('SetupWizardLandedSyncRepositories')
    }, [telemetryService])

    const handleFinishButtonClick = (): void => {
        telemetryService.log('SetupWizardFinishedSuccessfully')
    }

    return (
        <section {...props}>
            <Text className="mb-2">
                It may take a few moments to clone and index each repository. Repository statuses are displayed below.
            </Text>
            <SiteAdminRepositoriesContainer />

            <CustomNextButton label="Start searching" disabled={false} onClick={handleFinishButtonClick} />
        </section>
    )
}