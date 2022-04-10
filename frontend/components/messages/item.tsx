import React from 'react';
import LetterProfileIcon from '@app/components/icon/letter-profile-icon';
import {formatDate, formatFullString} from '@app/lib/date-format';

export interface Message {
    Date: string;
    MessageID: string;
    MessageKey: string;
    FromName: string;
    From: string;
    ThreadID: number;
    Subject: string;
    Tags: string[];
}

export interface MessageContent {
    Html: string;
    Text: string;
}

export interface ItemProps {
    message: Message;
    content: MessageContent;
}

export function Item({message, content}: ItemProps) {
    return (
        <li
            id={message.MessageID}
            data-key={message.MessageKey}
            data-thread={message.ThreadID.toString()}
            className="py-3 sm:py-4 mb-4 px-6 bg-white rounded-lg shadow-lg"
        >
            <div className="flex-1 min-w-0 flex justify-between mb-3">
                <p className="text-md">
                    <LetterProfileIcon name={message.FromName} />
                    <span className="ml-4">
                        {message.FromName} {'<' + message.From + '>'}
                    </span>
                </p>
                <p
                    className="text-md font-medium text-gray-900"
                    title={formatFullString(message.Date)}
                >
                    {formatDate(message.Date)}
                </p>
            </div>
            <div
                title={message.Subject}
                className="text-lg flex justify-between border-b pb-4"
            >
                <p className="truncate  font-medium text-gray-900 ">
                    {message.Subject}
                </p>
                <p className="font-medium text-gray-900">
                    ({message.Tags.join(', ')})
                </p>
            </div>
            <div className="whitespace-pre-line py-4">
                {content.Text.trim()}
            </div>
        </li>
    );
}

export default Item;
