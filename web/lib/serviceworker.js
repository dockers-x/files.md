const CACHE_NAME = 'files-md-v1';
const urlsToCache = [
    '/',
    '/favicon.ico',
    '/manifest.json',
    '/app.css?v=',
    '/lib/normalize.css?v=',
    '/lib/sidebar.css?v=',
    '/lib/codemirror.css?v=',
    '/lib/hypermd.css?v=',
    '/lib/theme-light.css?v=',
    '/lib/theme-dark.css?v=',
    '/chat.css?v=',
    '/lib/sidebar.js?v=',
    '/lib/codemirror.js?v=',
    '/lib/core.js?v=',
    '/lib/markdown.js?v=',
    '/lib/hypermd.js?v=',
    '/lib/keymap.js?v=',
    '/lib/click.js?v=',
    '/lib/hide-token.js?v=',
    '/lib/fold.js?v=',
    '/lib/fold-image.js?v=',
    '/lib/fold-link.js?v=',
    '/lib/table-align.js?v=',
    '/lib/autocomplete-link.js?v=',
    '/lib/show-hint.js?v=',
    '/lib/autoscroll.js?v=',
    '/lib/codemirror-go.js?v=',
    '/lib/codemirror-php.js?v=',
    '/lib/codemirror-shell.js?v=',
    '/lib/similarity.js?v=',
    '/welcome.js?v=',
    '/files.js?v=',
    '/wasm_exec.js?v=',
    '/app.js?v=',
    '/wasm.js?v=',
    '/chat.js?v=',
    '/modals.js?v=',
];

self.addEventListener('install', event => {
    event.waitUntil(
        caches.open(CACHE_NAME)
            .then(cache => {
                // Cache each file individually to find the problem
                const cachePromises = urlsToCache.map(url => {
                    console.log('Trying to cache:', url);
                    return cache.add(url)
                        .then(() => console.log('✓ Cached:', url))
                        .catch(err => console.error('✗ Failed to cache:', url, err));
                });
                return Promise.allSettled(cachePromises); // Won't fail if one fails
            })
    );
});

self.addEventListener('fetch', event => {
    console.log('intercepting');
    event.respondWith(
        caches.match(event.request)
            .then(response => {
                return response || fetch(event.request);
            })
    );
});