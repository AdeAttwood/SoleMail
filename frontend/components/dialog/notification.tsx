import React from 'react';
import ReactDOM from 'react-dom';
import ExclamationMark from '@app/components/icon/exclamation-mark';
import {error} from 'console';

const NOTIFY_ID = 'notify-container-' + (crypto as any).randomUUID();

export interface NotifyOptions {
    type: string;
    content: string;
    timeout?: number;
}

export interface NotificationProps extends NotifyOptions {
    id: string;
    color: string;
    icon: any;
}

interface NotificationState {
    notifications: {[key: string]: NotificationProps};
}

export const Notification = ({
    content,
    color,
    icon,
    ...rest
}: NotificationProps) => {
    const Icon = icon;
    return (
        <div
            {...rest}
            className="flex max-w-md bg-white shadow-lg rounded-lg overflow-hidden mb-4"
        >
            <div className={`w-2 bg-${color}`}></div>
            <div className="flex items-center px-2 py-3">
                <div className={`text-${color}`}>
                    <Icon />
                </div>
                <div className="mx-3">
                    <p className="text-gray-600">{content}</p>
                </div>
            </div>
        </div>
    );
};

class NotificationWrapper extends React.Component<{}, NotificationState> {
    states: any = {
        success: {
            classes: 'text-green-500 bg-green-500',
            color: 'green-500',
            icon: ExclamationMark,
        },
        error: {
            classes: 'text-red-500 bg-red-500',
            color: 'red-500',
            icon: ExclamationMark,
        },
    };

    constructor(props: any) {
        super(props);
        this.state = {notifications: {}};
    }

    add(options: NotifyOptions) {
        const {notifications} = this.state;
        const notification: NotificationProps = {
            ...options,
            ...this.states[options.type],
            id: 'notification-' + (crypto as any).randomUUID(),
        };

        setTimeout(() => this.remove(notification.id), options.timeout || 8000);

        notifications[notification.id] = notification;
        this.setState({notifications});
    }

    remove(id: string) {
        if (!this.state.notifications[id]) {
            return;
        }

        const {notifications} = this.state;
        delete notifications[id];
        this.setState({notifications});
    }

    render() {
        const notifications = this.state.notifications;
        const keys = Object.keys(notifications);

        return keys.map((id: string) => {
            return <Notification key={id} {...notifications[id]} />;
        });
    }
}

const notification_wrapper = React.createRef<NotificationWrapper>();

const getWrapper = async (): Promise<NotificationWrapper> => {
    if (!notification_wrapper.current) {
        let el = document.getElementById(NOTIFY_ID);
        if (!el) {
            el = document.createElement('div');
            el.id = NOTIFY_ID;
            el.classList.add('fixed', 'top-0', 'right-0', 'pt-6', 'pr-6');

            document.body.appendChild(el);
        }

        await new Promise<void>((resolve) => {
            ReactDOM.render(
                <NotificationWrapper ref={notification_wrapper} />,
                el,
                resolve,
            );
        });
    }

    if (!(notification_wrapper.current instanceof NotificationWrapper)) {
        throw new Error('Invlaid notification wrapper instance');
    }

    return notification_wrapper.current;
};

export function notify(options: NotifyOptions) {
    return getWrapper().then((wrapper) => wrapper.add(options));
}

notify.error = function (content: string) {
    return notify({type: 'error', content});
};

export default notify;
