import React from 'react';
import Button from '@app/components/core/button';
import loadView from '@app/lib/load-view';
// import * as ViewAccountCreate from "@app/views/account-create";

export interface WelcomeProps {}

export function Main() {
    return (
        <div className="max-w-md bg-white border-2 border-gray-300 p-6 rounded-md tracking-wide shadow-lg">
            <div className="flex flex-col items-center">
                <h1 className="mb-5">Sole Mails</h1>
                <Button
                    // onClick={() => loadView({ component: ViewAccountCreate })}
                    tooltip="This is some text"
                >
                    Get Started
                </Button>
            </div>
        </div>
    );
}
