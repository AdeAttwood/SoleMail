import React from 'react';
import {Message} from '@app/components/messages/item';
import {MessageList} from '@app/components/messages/list';
import promiseAllProperties from '@app/lib/promise-all-properties';
import IconBack from '@app/components/icon/back';
import IconArchive from '@app/components/icon/archive';
import IconTrash from '@app/components/icon/trash';
import loadView from '@app/lib/load-view';
import InputText from '@app/components/core/input-text';
import * as ThreadList from '@app/views/thread-list';

interface MainProps {
    threadId: number;
    messages: Message[];
    messageContent: {[key: string]: any};
}

export async function Props(props: any): Promise<MainProps> {
    if (typeof props.messages === 'undefined') {
        props.messages = await window.go.app.App.ReadThread(props.threadId);
    }

    props.messageContent = {};
    for (const message of props.messages) {
        console.log(message);
        props.messageContent[message.MessageID] =
            window.go.app.App.GetMessageContent(message.MessageID);
    }
    props.messageContent = await promiseAllProperties(props.messageContent);

    window.go.app.App.TagThread(props.threadId, '-unread');

    return props;
}

export function Main(props: MainProps) {
    const [tagString, setTagString] = React.useState<string>('');
    const tag = (tag_string: string) => {
        window.go.app.App.TagThread(props.threadId, tag_string).then(() =>
            loadView({component: ThreadList}),
        );
    };

    return (
        <div className="w-full block h-full">
            <div className="max-w-full flex justify-between mx-6 mt-6 p-4 bg-white rounded-lg shadow-lg">
                <div>
                    <IconBack
                        width={30}
                        height={30}
                        className="cursor-pointer inline-block mr-6"
                        onClick={() => {
                            loadView({component: ThreadList});
                        }}
                    />
                    <div className="inline-block" style={{width: '30px'}} />
                </div>
                <div style={{width: '600px'}}>
                    <form
                        onSubmit={(e) => {
                            e.preventDefault();
                            tag(tagString);
                        }}
                    >
                        <InputText
                            id="thread-tag"
                            value={tagString}
                            onChange={(e) => setTagString(e.target.value)}
                        />
                    </form>
                </div>
                <div>
                    <IconArchive
                        width={30}
                        height={30}
                        className="cursor-pointer inline-block mr-6"
                        onClick={() => tag('-inbox +archive')}
                    />
                    <IconTrash
                        width={30}
                        height={30}
                        className="cursor-pointer inline-block"
                        onClick={() => tag('-inbox +trash')}
                    />
                </div>
            </div>

            <div className="m-6">
                <MessageList
                    messages={props.messages}
                    contents={props.messageContent}
                />
            </div>
        </div>
    );
}
