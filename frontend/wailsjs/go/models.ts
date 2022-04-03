/* Do not change, this code is generated from Golang structs */

export {};

export class Config {
    static createFrom(source: any = {}) {
        return new Config(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
    }
}
export class EmailContent {
    static createFrom(source: any = {}) {
        return new EmailContent(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
    }
}
export class Message {
    static createFrom(source: any = {}) {
        return new Message(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
    }
}
export class Thread {
    static createFrom(source: any = {}) {
        return new Thread(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
    }
}
