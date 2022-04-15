import React from 'react';
import ThreadList from '@app/components/threads/list';
import {Thread} from '@app/components/threads/item';
import InputText from '@app/components/core/input-text';
import {useKeyboard} from '@app/hooks/use-keyboard';
import NumberIcon from '@app/components/icon/number-icon';

interface MainProps {
    threads: Thread[];
}

let QUERY: string | undefined = undefined;

export async function Props(props: any): Promise<MainProps> {
    if (typeof QUERY === 'undefined') {
        QUERY = 'tag:inbox';
    }

    if (typeof props.query !== 'undefined') {
        QUERY = props.query;
    }

    if (typeof props.threads === 'undefined') {
        props.threads = await window.go.app.App.GetThreads(
            QUERY || 'tag:inbox',
        );
    }

    return props;
}

type STATUS = 'idle' | 'loading';

export function Main({threads: initial_threads}: MainProps) {
    const [query, setQuery] = React.useState<string>(QUERY || '');
    const [status, setStatus] = React.useState<STATUS>('idle');
    const [threads, setThreads] = React.useState<Thread[]>(initial_threads);

    const search = (term: string) => {
        if (
            term.length &&
            ['+', '-'].includes(term.trim().at(0) || '') &&
            typeof QUERY !== 'undefined'
        ) {
            setStatus('loading');
            window.go.app.App.TagQuery(QUERY, term).then(() => {
                setQuery(QUERY || 'tag:inbox');
                search(QUERY || 'tag:inbox');
                setStatus('idle');
            });

            return;
        }

        setStatus('loading');
        window.go.app.App.GetThreads(term).then((t: any) => {
            if (t.length) {
                setThreads(t);
                setStatus('idle');
                QUERY = term;
            }
        });
    };

    const update = () => {
        setStatus('loading');
        window.go.app.App.Update().then(() => search(query));
    };

    useKeyboard(document, {
        '/': () => {
            const el = document.getElementById('thread-query');
            if (el !== null && document.activeElement !== el) {
                el.focus();
            }
        },
        'CTRL-r': () => update(),
    });

    React.useEffect(() => {
        const interval = setInterval(update, 300000);
        return () => clearInterval(interval);
    }, []);

    return (
        <div className="w-full block h-full">
            <div className="max-w-full flex justify-between mx-6 mt-6 p-4 bg-white rounded-lg shadow-lg">
                <div>
                    <NumberIcon
                        cursor="pointer"
                        width="60px"
                        number={threads.length}
                        spinning={status === 'loading'}
                        onClick={() => update()}
                    />
                </div>
                <div style={{width: '600px'}}>
                    <form
                        onSubmit={(e) => {
                            e.preventDefault();
                            search(query);
                        }}
                    >
                        <InputText
                            id="thread-query"
                            value={query}
                            onChange={(e) => {
                                setQuery(e.target.value);
                            }}
                        />
                    </form>
                </div>
                <div></div>
            </div>

            <div
                className="scroll-container m-6 bg-white rounded-lg shadow-lg"
                style={{maxHeight: 'calc(100vh - 142px)', overflowY: 'auto'}}
            >
                <ThreadList threads={threads.slice(0, 100)} />
            </div>
        </div>
    );
}
