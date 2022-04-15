import React from 'react';
import {Thread, Item} from './item';

export interface ThreadListProps {
    threads: Thread[];
}

export function ThreadList({threads}: ThreadListProps) {
    return (
        <ul
            role="list"
            className="divide-y divide-gray-200 dark:divide-gray-700 bg-white rounded-lg shadow-lg cursor-pointer"
        >
            {threads.map((thread) => (
                <Item key={thread.ThreadID} thread={thread} />
            ))}
        </ul>
    );
}

export default ThreadList;
