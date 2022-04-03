export interface go {
    app: {
        App: {
            GetConfig(): Promise<Config>;
            GetMessageContent(arg1: string): Promise<EmailContent | Error>;
            GetMessages(): Promise<Array<Message> | Error>;
            GetThreads(arg1: string): Promise<Array<Thread> | Error>;
            ReadThread(arg1: number): Promise<Array<Message> | Error>;
            TagThread(arg1: number, arg2: string): Promise<Thread | Error>;
        };
    };
}

declare global {
    interface Window {
        go: go;
    }
}
