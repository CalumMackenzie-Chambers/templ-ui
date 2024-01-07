import { WebSocketServer } from 'ws';
import { watch } from 'chokidar';
import pkg from 'lodash';

const { debounce } = pkg;

const wss = new WebSocketServer({ port: 1234 });

const sendReload = debounce(() => {
  console.log('Changes stabilized. Sending reload signal to browser...');
  wss.clients.forEach((client) => {
    console.log('Sending reload signal to browser...')
    client.send('reload');
  });
}, 300);

const watcher = watch(['./tmp/main', './docs/static/assets/**/*', './docs/static/css/styles.css', './docs/static/js/main.js'], {
  persistent: true,
  ignoreInitial: true,
  usePolling: false,
  depth: 99,
});

watcher.on('change', () => {
  console.log('Changes detected. Waiting for stabilization...');
  sendReload();
});
