import * as React from 'react';
import {ReactNode} from 'react';
import {Button} from './button';
import {Icon} from './icon';

export const ToggleButton = ({
    title,
    children,
    onToggle,
    toggled,
    disabled,
    icon
}: {
    toggled: boolean;
    onToggle: () => void;
    children?: ReactNode;
    title: string;
    disabled?: boolean;
    icon: Icon;
}) => (
    <Button title={title} onClick={onToggle} icon={icon} disabled={disabled}>
        {children}
    </Button>
);
