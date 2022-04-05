import React from 'react';

export type KeyMap = {[key: string]: (event: KeyboardEvent) => void};
export type EventSource = Document;

export const useKeyboard = (element: EventSource, keyMap: KeyMap) => {
    React.useEffect(() => {
        const onKeyup = (event: KeyboardEvent) => {
            let key = event.key;
            if (event.ctrlKey) {
                key = `CRTL-${key}`;
            }

            if (typeof keyMap[key] !== 'undefined') {
                keyMap[key](event);
            }
        };

        element.addEventListener('keyup', onKeyup, false);
        return () => element.removeEventListener('keyup', onKeyup, false);
    }, []);
};
