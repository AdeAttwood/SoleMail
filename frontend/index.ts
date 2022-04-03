import loadView from '@app/lib/load-view';

// import the styles
import '@app/index.css';

// import the views
import * as Welcome from '@app/views/welcome';
import * as ThreadList from '@app/views/thread-list';

function boot(threads: any) {
    if (threads.length) {
        loadView({component: ThreadList, props: {threads}});
    } else {
        loadView({component: Welcome});
    }
}

window.go.app.App.GetThreads('tag:inbox').then(boot);
