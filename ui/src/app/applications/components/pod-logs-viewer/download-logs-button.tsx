import {services} from '../../../shared/services';
import * as React from 'react';
import {PodLogsProps} from './pod-logs-viewer';
import {Button} from '../../../shared/components/button';

export const DownloadLogsButton = ({applicationName, applicationNamespace, containerName, group, kind, name, namespace, podName}: PodLogsProps) => (
    <Button
        title='Download logs'
        icon='download'
        onClick={async () => {
            const downloadURL = services.applications.getDownloadLogsURL(applicationName, applicationNamespace, namespace, podName, {group, kind, name}, containerName);
            window.open(downloadURL, '_blank');
        }}
    />
);
