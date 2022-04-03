import React from 'react';

interface LetterIconProps {
    name: string;
}

export const LetterProfileIcon = ({name}: LetterIconProps) => {
    let text = 'AA';
    const letters = /(\w{1})([^ ]+ (\w))?/.exec(name);
    if (letters !== null) {
        text = `${letters[1]}${letters[3] || ''}`;
    }

    return (
        <span className="inline-block">
            <span
                className="flex justify-center items-center bg-red-700"
                style={{width: '35px', height: '35px', borderRadius: '50%'}}
            >
                <span className="tracking-widest text-white font-bold">
                    {text}
                </span>
            </span>
        </span>
    );
};

export default LetterProfileIcon;
