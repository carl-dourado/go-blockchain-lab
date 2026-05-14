package server

const explorerHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Go Blockchain Lab</title>
  <style>
    :root {
      color-scheme: dark;
      --bg: #08111f;
      --panel: #101a2c;
      --panel-strong: #16233a;
      --text: #e5edf8;
      --muted: #8ba0ba;
      --line: #243553;
      --accent: #00add8;
      --ok: #32d583;
      --warn: #fdb022;
      --bad: #f97066;
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace;
    }

    * {
      box-sizing: border-box;
    }

    body {
      margin: 0;
      background: radial-gradient(circle at 50% -10%, rgba(0, 173, 216, 0.22), transparent 34rem), var(--bg);
      color: var(--text);
    }

    button,
    input,
    textarea {
      font: inherit;
    }

    .shell {
      width: min(1120px, calc(100% - 32px));
      margin: 0 auto;
      padding: 42px 0;
    }

    header {
      display: flex;
      justify-content: space-between;
      gap: 18px;
      align-items: end;
      margin-bottom: 24px;
    }

    h1,
    h2,
    h3,
    p {
      margin: 0;
    }

    .eyebrow {
      color: var(--accent);
      font-size: 12px;
      letter-spacing: 0;
      text-transform: uppercase;
    }

    h1 {
      margin-top: 8px;
      font-size: clamp(30px, 6vw, 58px);
      line-height: 0.95;
    }

    .toolbar {
      display: flex;
      gap: 10px;
      flex-wrap: wrap;
      justify-content: flex-end;
    }

    .grid {
      display: grid;
      grid-template-columns: repeat(4, minmax(0, 1fr));
      gap: 14px;
      margin-bottom: 14px;
    }

    .panel,
    .stat,
    .block,
    .record {
      border: 1px solid var(--line);
      background: rgba(16, 26, 44, 0.88);
      border-radius: 8px;
    }

    .stat {
      min-height: 92px;
      padding: 14px;
    }

    .stat span {
      display: block;
      color: var(--muted);
      font-size: 12px;
      margin-bottom: 10px;
    }

    .stat strong {
      font-size: 26px;
      color: var(--text);
      overflow-wrap: anywhere;
    }

    .panel {
      padding: 18px;
      margin-bottom: 14px;
    }

    .panel-head {
      display: flex;
      justify-content: space-between;
      align-items: center;
      gap: 12px;
      margin-bottom: 14px;
    }

    .panel-head p {
      color: var(--muted);
      font-size: 13px;
      margin-top: 4px;
    }

    form {
      display: grid;
      grid-template-columns: minmax(140px, 220px) 1fr auto;
      gap: 10px;
      align-items: end;
    }

    label {
      display: grid;
      gap: 6px;
      color: var(--muted);
      font-size: 12px;
    }

    input,
    textarea {
      width: 100%;
      border: 1px solid var(--line);
      border-radius: 8px;
      background: #07101e;
      color: var(--text);
      padding: 12px;
      outline: none;
    }

    textarea {
      min-height: 46px;
      resize: vertical;
    }

    button {
      border: 1px solid rgba(0, 173, 216, 0.42);
      border-radius: 8px;
      color: #00121a;
      background: var(--accent);
      padding: 12px 14px;
      cursor: pointer;
      min-height: 44px;
    }

    button.secondary {
      background: transparent;
      color: var(--text);
      border-color: var(--line);
    }

    button:disabled {
      opacity: 0.55;
      cursor: wait;
    }

    .status {
      min-height: 24px;
      color: var(--muted);
      font-size: 13px;
      margin-top: 12px;
    }

    .status.ok {
      color: var(--ok);
    }

    .status.bad {
      color: var(--bad);
    }

    .layout {
      display: grid;
      grid-template-columns: 1.8fr 1fr;
      gap: 14px;
    }

    .list {
      display: grid;
      gap: 12px;
    }

    .block {
      padding: 16px;
    }

    .block-top,
    .record {
      display: grid;
      gap: 6px;
    }

    .block-top {
      grid-template-columns: auto 1fr auto;
      align-items: center;
      gap: 12px;
      margin-bottom: 12px;
    }

    .height {
      width: 44px;
      height: 44px;
      display: grid;
      place-items: center;
      border-radius: 8px;
      background: var(--panel-strong);
      color: var(--accent);
      font-weight: 800;
    }

    .hash,
    .muted {
      color: var(--muted);
      overflow-wrap: anywhere;
    }

    .hash {
      font-size: 12px;
    }

    .pill {
      display: inline-flex;
      align-items: center;
      border-radius: 999px;
      border: 1px solid var(--line);
      color: var(--muted);
      padding: 6px 8px;
      font-size: 12px;
      white-space: nowrap;
    }

    .records {
      display: grid;
      gap: 8px;
      margin-top: 10px;
    }

    .record {
      padding: 10px;
      background: #0b1526;
    }

    .record b {
      color: var(--text);
      overflow-wrap: anywhere;
    }

    .empty {
      color: var(--muted);
      border: 1px dashed var(--line);
      border-radius: 8px;
      padding: 18px;
      text-align: center;
    }

    @media (max-width: 820px) {
      header,
      .layout {
        display: grid;
      }

      .toolbar {
        justify-content: start;
      }

      .grid,
      form,
      .layout {
        grid-template-columns: 1fr;
      }

      .block-top {
        grid-template-columns: auto 1fr;
      }

      .block-top .pill {
        grid-column: 1 / -1;
      }
    }
  </style>
</head>
<body>
  <main class="shell">
    <header>
      <div>
        <p class="eyebrow">local blockchain explorer</p>
        <h1>Go Blockchain Lab</h1>
      </div>
      <div class="toolbar">
        <button class="secondary" id="refresh" type="button">refresh</button>
        <button id="mine" type="button">mine pending</button>
      </div>
    </header>

    <section class="grid" aria-label="chain summary">
      <article class="stat"><span>height</span><strong id="height">-</strong></article>
      <article class="stat"><span>pending</span><strong id="pending-count">-</strong></article>
      <article class="stat"><span>valid</span><strong id="valid">-</strong></article>
      <article class="stat"><span>latest hash</span><strong class="hash" id="latest-hash">-</strong></article>
    </section>

    <section class="panel">
      <div class="panel-head">
        <div>
          <h2>Add record</h2>
          <p>Records stay pending until a new block is mined.</p>
        </div>
      </div>
      <form id="record-form">
        <label>
          author
          <input id="author" value="carlos" autocomplete="off">
        </label>
        <label>
          data
          <textarea id="data">transfer 10 from alice to bob</textarea>
        </label>
        <button type="submit">add</button>
      </form>
      <p class="status" id="status"></p>
    </section>

    <section class="layout">
      <section class="panel">
        <div class="panel-head">
          <div>
            <h2>Blocks</h2>
            <p>Each block links to the previous hash.</p>
          </div>
        </div>
        <div class="list" id="blocks"></div>
      </section>

      <aside class="panel">
        <div class="panel-head">
          <div>
            <h2>Pending</h2>
            <p>These records are not in a block yet.</p>
          </div>
        </div>
        <div class="list" id="pending"></div>
      </aside>
    </section>
  </main>

  <script>
    const state = {
      busy: false
    };

    const $ = (id) => document.getElementById(id);
    const shortHash = (hash) => hash ? hash.slice(0, 14) + '...' + hash.slice(-8) : '-';
    const formatDate = (value) => value ? new Date(value).toLocaleString() : '-';

    function setStatus(message, tone = '') {
      const node = $('status');
      node.textContent = message;
      node.className = tone ? 'status ' + tone : 'status';
    }

    function setBusy(isBusy) {
      state.busy = isBusy;
      document.querySelectorAll('button, input, textarea').forEach((node) => {
        node.disabled = isBusy;
      });
    }

    async function request(path, options = {}) {
      const response = await fetch(path, {
        headers: { 'Content-Type': 'application/json' },
        ...options
      });
      const payload = await response.json();
      if (!response.ok) {
        throw new Error(payload.error || response.statusText);
      }
      return payload;
    }

    function renderRecord(record) {
      const item = document.createElement('article');
      item.className = 'record';

      const title = document.createElement('b');
      title.textContent = record.data;
      item.appendChild(title);

      const meta = document.createElement('span');
      meta.className = 'muted';
      meta.textContent = record.author + ' | ' + formatDate(record.createdAt) + ' | ' + record.id;
      item.appendChild(meta);

      return item;
    }

    function renderBlocks(blocks) {
      const list = $('blocks');
      list.replaceChildren();

      if (!blocks.length) {
        const empty = document.createElement('p');
        empty.className = 'empty';
        empty.textContent = 'no blocks';
        list.appendChild(empty);
        return;
      }

      [...blocks].reverse().forEach((block) => {
        const item = document.createElement('article');
        item.className = 'block';

        const top = document.createElement('div');
        top.className = 'block-top';

        const height = document.createElement('span');
        height.className = 'height';
        height.textContent = block.index;
        top.appendChild(height);

        const text = document.createElement('div');
        const title = document.createElement('h3');
        title.textContent = block.index === 0 ? 'genesis block' : 'block #' + block.index;
        text.appendChild(title);

        const hash = document.createElement('p');
        hash.className = 'hash';
        hash.textContent = 'hash ' + block.hash;
        text.appendChild(hash);
        top.appendChild(text);

        const nonce = document.createElement('span');
        nonce.className = 'pill';
        nonce.textContent = 'nonce ' + block.nonce;
        top.appendChild(nonce);

        item.appendChild(top);

        const prev = document.createElement('p');
        prev.className = 'hash';
        prev.textContent = 'prev ' + block.previousHash;
        item.appendChild(prev);

        const meta = document.createElement('p');
        meta.className = 'muted';
        meta.textContent = formatDate(block.timestamp) + ' | difficulty ' + block.difficulty;
        item.appendChild(meta);

        const records = document.createElement('div');
        records.className = 'records';
        block.records.forEach((record) => records.appendChild(renderRecord(record)));
        item.appendChild(records);

        list.appendChild(item);
      });
    }

    function renderPending(records) {
      const list = $('pending');
      list.replaceChildren();

      if (!records.length) {
        const empty = document.createElement('p');
        empty.className = 'empty';
        empty.textContent = 'no pending records';
        list.appendChild(empty);
        return;
      }

      records.forEach((record) => list.appendChild(renderRecord(record)));
    }

    async function load() {
      const [chain, pending, validation] = await Promise.all([
        request('/chain'),
        request('/pending'),
        request('/validate')
      ]);
      const latest = chain[chain.length - 1];

      $('height').textContent = chain.length ? chain.length - 1 : 0;
      $('pending-count').textContent = pending.length;
      $('valid').textContent = validation.valid ? 'yes' : 'no';
      $('valid').style.color = validation.valid ? 'var(--ok)' : 'var(--bad)';
      $('latest-hash').textContent = shortHash(latest && latest.hash);

      renderBlocks(chain);
      renderPending(pending);
    }

    async function run(action, success) {
      if (state.busy) return;
      setBusy(true);
      setStatus('');
      try {
        await action();
        await load();
        setStatus(success, 'ok');
      } catch (error) {
        setStatus(error.message, 'bad');
      } finally {
        setBusy(false);
      }
    }

    $('record-form').addEventListener('submit', (event) => {
      event.preventDefault();
      run(
        () => request('/records', {
          method: 'POST',
          body: JSON.stringify({
            author: $('author').value,
            data: $('data').value
          })
        }),
        'record added to pending queue'
      );
    });

    $('mine').addEventListener('click', () => {
      run(() => request('/mine', { method: 'POST' }), 'block mined');
    });

    $('refresh').addEventListener('click', () => {
      run(() => Promise.resolve(), 'refreshed');
    });

    load().catch((error) => setStatus(error.message, 'bad'));
  </script>
</body>
</html>
`
