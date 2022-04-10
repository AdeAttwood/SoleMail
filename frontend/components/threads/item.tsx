import React from 'react';
import LetterProfileIcon from '@app/components/icon/letter-profile-icon';
import loadView from '@app/lib/load-view';
import * as ThreadList from '@app/views/thread-view';
import {formatDate, formatFullString} from '@app/lib/date-format';

export interface Thread {
    Date: string;
    FromName: string;
    From: string;
    ThreadID: number;
    Subject: string;
    Messages: string[];
    Tags: string[];
}

export interface ItemProps {
    thread: Thread;
}

export function Item({thread}: ItemProps) {
    const onClick = () => {
        loadView({
            component: ThreadList,
            props: {threadId: thread.ThreadID},
        });
    };

    return (
        <li className="py-3" onClick={onClick} id={`t_${thread.ThreadID}`}>
            <div className="flex-1 min-w-0 flex justify-between mb-2">
                <p className="text-md">
                    <LetterProfileIcon name={thread.FromName} />
                    <span className="ml-4">
                        {thread.FromName} {'<' + thread.From + '>'}
                    </span>
                </p>
                <p
                    className="text-md font-medium text-gray-900"
                    title={formatFullString(thread.Date)}
                >
                    {formatDate(thread.Date)}
                </p>
            </div>
            <div
                title={thread.Subject}
                className="text-lg flex justify-between"
            >
                <p className="truncate flex-1 font-medium text-gray-900">
                    {thread.Subject}
                </p>
                <p className="font-medium text-gray-900">
                    ({thread.Tags.join(', ')}) ({thread.Messages.length})
                </p>
            </div>
        </li>
    );
}

export default Item;
