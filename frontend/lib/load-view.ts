import React from 'react';
import ReactDOM from 'react-dom';

export interface ViewComponent<T> {
    Main: React.FC<T>;
    Props?: (props: any) => Promise<T>;
}

export interface LoadViewProps<T> {
    component: ViewComponent<T>;
    props?: any;
}

export async function loadView<T>(args: LoadViewProps<T>) {
    const component = args.component;
    let props = args.props || {};
    if (typeof component.Props === 'function') {
        props = {
            ...(await component.Props(props)),
            ...props,
        };
    }

    ReactDOM.render(
        React.createElement(component.Main, props),
        document.getElementById('root'),
    );
}

export default loadView;
