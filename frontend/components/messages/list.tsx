import React from 'react';
import {Message, Item} from './item';

export interface ThreadListProps {
    messages: Message[];
    contents: {[key: string]: any};
}

export function MessageList({messages, contents}: ThreadListProps) {
    return (
        <ul role="list" className="mb-4">
            {messages.map((message) => (
                <Item
                    key={message.MessageID}
                    message={message}
                    content={contents[message.MessageID]}
                />
            ))}
        </ul>
    );
}

export default MessageList;
