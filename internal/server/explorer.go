package server

const explorerHTML = `<!doctype html>
<html lang="pt-BR">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Go Blockchain Lab</title>
  <style>
    :root {
      color-scheme: dark;
      --bg: #020707;
      --screen: #171717;
      --text: #f4f4f4;
      --muted: #a7b4b4;
      --line: #77d7e8;
      --blue: #1499ff;
      --yellow: #ffff00;
      --red: #ff2020;
      --ok: #65ff8f;
      --bad: #ff4d4d;
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace;
    }

    * {
      box-sizing: border-box;
    }

    body {
      margin: 0;
      min-height: 100vh;
      background: var(--bg);
      color: var(--text);
      padding: 5px;
    }

    button,
    input,
    textarea {
      font: inherit;
    }

    .shell {
      width: min(1180px, 100%);
      min-height: calc(100vh - 10px);
      margin: 0 auto;
      padding: 24px;
      background: var(--screen);
    }

    .terminal-menu {
      display: flex;
      justify-content: space-between;
      gap: 18px;
      color: var(--text);
      font-size: clamp(13px, 2vw, 16px);
      line-height: 1.55;
      margin-bottom: 28px;
    }

    .terminal-menu-left,
    .terminal-clock {
      display: grid;
      gap: 2px;
    }

    .terminal-menu-row {
      display: flex;
      gap: clamp(20px, 5vw, 34px);
    }

    .terminal-clock span {
      display: grid;
      grid-template-columns: 50px 1fr;
      gap: 10px;
    }

    .ascii-stage {
      display: grid;
      grid-template-columns: minmax(150px, 0.34fr) minmax(0, 1fr);
      gap: clamp(20px, 5vw, 54px);
      align-items: end;
      margin-bottom: 18px;
    }

    .ascii-cat,
    .ascii-logo,
    .ascii-output {
      margin: 0;
      white-space: pre;
      line-height: 1.16;
    }

    .ascii-cat {
      color: var(--text);
      font-size: clamp(9px, 1vw, 13px);
    }

    .ascii-logo {
      font-weight: 800;
      font-size: clamp(8px, 1vw, 13px);
    }

    .ascii-red {
      color: var(--red);
    }

    .ascii-white {
      color: var(--text);
    }

    .ascii-blue {
      color: var(--blue);
    }

    .terminal-subtitle {
      display: grid;
      gap: 8px;
      margin-left: clamp(12px, 7vw, 70px);
      margin-bottom: 20px;
      line-height: 1.35;
    }

    .terminal-title {
      color: var(--yellow);
      font-size: clamp(18px, 2.5vw, 24px);
      font-weight: 800;
    }

    .terminal-yellow {
      color: var(--yellow);
      font-weight: 800;
    }

    .terminal-blue {
      color: var(--blue);
    }

    .ascii-output {
      width: 100%;
      color: var(--text);
      background: transparent;
      font-size: clamp(11px, 1.42vw, 15px);
      margin-bottom: 16px;
      overflow-x: auto;
      overflow-y: hidden;
    }

    .ascii-output.is-blue {
      color: var(--blue);
    }

    .ascii-output.is-yellow {
      color: var(--yellow);
    }

    .terminal-form {
      display: grid;
      gap: 12px;
      margin-bottom: 16px;
    }

    .terminal-form form {
      display: grid;
      gap: 10px;
    }

    .terminal-field {
      display: grid;
      grid-template-columns: 110px minmax(0, 1fr);
      gap: 10px;
      align-items: start;
      color: var(--yellow);
    }

    input,
    textarea {
      width: 100%;
      color: var(--text);
      background: #050505;
      border: 0;
      border-bottom: 1px solid var(--line);
      outline: none;
      padding: 2px 4px 4px;
    }

    textarea {
      min-height: 44px;
      resize: vertical;
    }

    .terminal-actions {
      display: flex;
      gap: 12px;
      flex-wrap: wrap;
      align-items: center;
    }

    button {
      min-height: 34px;
      color: var(--yellow);
      background: transparent;
      border: 0;
      border-bottom: 1px solid var(--line);
      padding: 4px 2px;
      cursor: pointer;
    }

    button:disabled {
      color: var(--muted);
      cursor: wait;
    }

    .status {
      min-height: 22px;
      color: var(--yellow);
    }

    .status.ok {
      color: var(--ok);
    }

    .status.bad {
      color: var(--bad);
    }

    .ascii-columns {
      display: grid;
      grid-template-columns: 1.45fr 0.9fr;
      gap: 20px;
    }

    .terminal-status {
      display: flex;
      justify-content: space-between;
      gap: 18px;
      margin-top: 18px;
      color: var(--text);
      font-size: 15px;
    }

    .terminal-status span {
      white-space: pre;
    }

    .prompt-cursor {
      display: inline-block;
      width: 10px;
      height: 1.1em;
      margin-left: 6px;
      border-left: 2px solid var(--line);
      transform: translateY(3px);
      animation: blink 1s steps(2, start) infinite;
    }

    @keyframes blink {
      50% {
        opacity: 0;
      }
    }

    @media (max-width: 820px) {
      .terminal-menu,
      .ascii-stage,
      .ascii-columns,
      .terminal-status {
        display: grid;
      }

      .terminal-field {
        grid-template-columns: 1fr;
      }

      .ascii-logo {
        font-size: clamp(10px, 4vw, 18px);
      }
    }
  </style>
</head>
<body>
  <main class="shell">
    <header>
      <div class="terminal-menu">
        <div class="terminal-menu-left">
          <div class="terminal-menu-row">
            <span>File</span>
            <span>Options</span>
            <span>Keypad</span>
          </div>
          <div class="terminal-menu-row">
            <span>Terminal</span>
            <span>GOCHAIN</span>
          </div>
        </div>
        <div class="terminal-clock">
          <span>Date <b id="date-value">--.--.--</b></span>
          <span>Time <b id="time-value">--:--:--</b></span>
        </div>
      </div>

      <div class="ascii-stage">
        <pre class="ascii-cat">      |\      _,,,---,,_
ZZZzz /,.-''      -.  ;-;;
     |,4-  ) )-,_. ,\ (  '-'
    '---''(_/--'  '-'\_)</pre>
        <pre class="ascii-logo"><span class="ascii-red">   ______   ______</span>   <span class="ascii-white"> ______   __  __</span>   <span class="ascii-blue"> ______   ______   __   __</span>
<span class="ascii-red">  / ____/  / __  /</span>   <span class="ascii-white">/ ____/  / / / /</span>   <span class="ascii-blue">/ ____/  / __  /  /  | / /</span>
<span class="ascii-red"> / / __   / / / /</span>   <span class="ascii-white">/ /      / /_/ /</span>   <span class="ascii-blue">/ /      / /_/ /  / /| |/ /</span>
<span class="ascii-red">/ /_/ /  / /_/ /</span>   <span class="ascii-white">/ /___   / __  /</span>   <span class="ascii-blue">/ /___   / __  /  / / |   /</span>
<span class="ascii-red">\____/   \____/</span>   <span class="ascii-white">\____/  /_/ /_/</span>   <span class="ascii-blue">\____/  /_/ /_/  /_/  |__/</span></pre>
      </div>

      <div class="terminal-subtitle">
        <span class="terminal-title">Go Blockchain Lab</span>
        <span class="terminal-yellow">The GOCHAIN 0.1 local blockchain system</span>
        <span class="terminal-blue">records -> pending queue -> mine block -> validate previous hash</span>
      </div>
    </header>

    <pre id="chain-art" class="ascii-output is-blue">+----------------------------+
| loading ASCII blockchain  |
+----------------------------+</pre>

    <pre id="summary-art" class="ascii-output is-yellow">+----------------+
| loading status |
+----------------+</pre>

    <section class="terminal-form" aria-label="Comandos do lab">
      <pre class="ascii-output">+--------------------------- COMMAND INPUT ---------------------------+
| 1. add record     2. mine pending records     3. validate chain    |
+---------------------------------------------------------------------+</pre>
      <form id="record-form">
        <label class="terminal-field">
          <span>AUTHOR ===&gt;</span>
          <input id="author" value="carlos" autocomplete="off">
        </label>
        <label class="terminal-field">
          <span>DATA   ===&gt;</span>
          <textarea id="data">transfer 10 from alice to bob</textarea>
        </label>
        <div class="terminal-actions">
          <button type="submit">PF1 ADD-RECORD</button>
          <button id="mine" type="button">PF2 MINE-BLOCK</button>
          <button id="validate-now" type="button">PF3 VALIDATE</button>
          <button id="refresh" type="button">PF5 REFRESH</button>
        </div>
      </form>
      <p class="status" id="status"></p>
    </section>

    <section class="ascii-columns">
      <pre id="blocks" class="ascii-output">+----------+
| BLOCKS  |
+----------+</pre>
      <pre id="pending" class="ascii-output">+----------+
| PENDING |
+----------+</pre>
    </section>

    <footer class="terminal-status">
      <span>Logon ===&gt; <b id="prompt-status">READY</b><i class="prompt-cursor" aria-hidden="true"></i></span>
      <span>RUNNING  GOCHAIN</span>
    </footer>
  </main>

  <script>
    const state = {
      busy: false
    };

    const $ = (id) => document.getElementById(id);
    const shortHash = (hash) => hash ? hash.slice(0, 14) + '...' + hash.slice(-8) : '-';
    const tinyHash = (hash) => hash ? hash.slice(0, 8) : '--------';
    const formatDate = (value) => value ? new Date(value).toLocaleString() : '-';

    function pad(value) {
      return String(value).padStart(2, '0');
    }

    function fit(value, width) {
      const text = String(value ?? '');
      if (text.length > width) return text.slice(0, Math.max(0, width - 1)) + '~';
      return text.padEnd(width, ' ');
    }

    function line(width) {
      return '+' + '-'.repeat(width - 2) + '+';
    }

    function box(title, rows, width = 70) {
      const inner = width - 4;
      const out = [line(width), '| ' + fit(title, inner) + ' |', line(width)];
      rows.forEach((row) => out.push('| ' + fit(row, inner) + ' |'));
      out.push(line(width));
      return out.join('\n');
    }

    function updateClock() {
      const now = new Date();
      $('date-value').textContent = pad(now.getDate()) + '.' + pad(now.getMonth() + 1) + '.' + String(now.getFullYear()).slice(-2);
      $('time-value').textContent = pad(now.getHours()) + ':' + pad(now.getMinutes()) + ':' + pad(now.getSeconds());
    }

    function setStatus(message, tone = '') {
      const node = $('status');
      node.textContent = message;
      node.className = tone ? 'status ' + tone : 'status';
      $('prompt-status').textContent = tone === 'bad' ? 'ERROR' : message ? 'OK' : 'READY';
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

    function cubeLines(block) {
      const index = '#' + String(block.index).padStart(3, '0');
      const hash = tinyHash(block.hash);
      const recs = 'rec ' + String(block.records?.length || 0).padStart(2, '0');
      return [
        '   .--------.   ',
        '  / ' + fit(index, 5) + ' /|  ',
        ' +--------+ |  ',
        ' | ' + fit(hash, 6) + ' | +  ',
        ' | ' + fit(recs, 6) + ' |/   ',
        ' +--------+    '
      ];
    }

    function renderChainArt(blocks, validation) {
      const visible = blocks.length > 5 ? blocks.slice(-5) : blocks;
      const status = validation.valid ? 'CHAIN VALID' : 'CHAIN BROKEN';
      const rows = [
        '+---------------------------- GOCHAIN ASCII VIEW -----------------------------+',
        '| cada cubo e um bloco; cada ==o==o== e um elo pelo previousHash             |',
        '| status: ' + fit(status, 66) + '|',
        '+----------------------------------------------------------------------------+',
        ''
      ];

      if (!visible.length) {
        rows.push('                         [ nenhum bloco ainda ]');
        return rows.join('\n');
      }

      const cubes = visible.map(cubeLines);
      for (let lineIndex = 0; lineIndex < cubes[0].length; lineIndex++) {
        let row = '';
        cubes.forEach((cube, index) => {
          row += cube[lineIndex];
          if (index < cubes.length - 1) {
            row += lineIndex === 2 || lineIndex === 3 ? '==o==o==>' : '         ';
          }
        });
        rows.push(row);
      }

      if (blocks.length > visible.length) {
        rows.push('');
        rows.push('mostrando os ultimos ' + visible.length + ' de ' + blocks.length + ' blocos');
      }

      return rows.join('\n');
    }

    function renderSummary(chain, pending, validation) {
      const latest = chain[chain.length - 1];
      return box('CURRENT STATE', [
        'height........: ' + (chain.length ? chain.length - 1 : 0),
        'pending.......: ' + pending.length,
        'validation....: ' + (validation.valid ? 'OK' : 'FAILED'),
        'latest hash...: ' + shortHash(latest && latest.hash),
        'next action...: ' + (pending.length ? 'PF2 MINE-BLOCK' : 'PF1 ADD-RECORD')
      ], 80);
    }

    function renderBlocks(blocks) {
      if (!blocks.length) {
        $('blocks').textContent = box('BLOCKS', ['nenhum bloco encontrado'], 78);
        return;
      }

      const rows = [];
      [...blocks].reverse().forEach((block) => {
        rows.push('BLOCK #' + block.index + '  nonce=' + block.nonce + '  difficulty=' + block.difficulty);
        rows.push('  hash : ' + block.hash);
        rows.push('  prev : ' + block.previousHash);
        rows.push('  time : ' + formatDate(block.timestamp));
        rows.push('  records: ' + block.records.length);
        block.records.slice(0, 4).forEach((record) => {
          rows.push('    - ' + record.author + ' :: ' + record.data);
        });
        if (block.records.length > 4) rows.push('    ... +' + (block.records.length - 4) + ' records');
        rows.push('');
      });

      $('blocks').textContent = box('MINED BLOCKS', rows, 94);
    }

    function renderPending(records) {
      if (!records.length) {
        $('pending').textContent = box('PENDING QUEUE', ['sem registros pendentes'], 58);
        return;
      }

      const rows = records.map((record, index) => {
        return String(index + 1).padStart(2, '0') + '. ' + record.author + ' :: ' + record.data;
      });
      $('pending').textContent = box('PENDING QUEUE', rows, 58);
    }

    async function load() {
      const [chain, pending, validation] = await Promise.all([
        request('/chain'),
        request('/pending'),
        request('/validate')
      ]);

      $('chain-art').textContent = renderChainArt(chain, validation);
      $('summary-art').textContent = renderSummary(chain, pending, validation);
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
        'registro entrou na fila pendente'
      );
    });

    $('mine').addEventListener('click', () => {
      run(() => request('/mine', { method: 'POST' }), 'bloco minerado e ligado na corrente');
    });

    $('refresh').addEventListener('click', () => {
      run(() => Promise.resolve(), 'estado atualizado');
    });

    $('validate-now').addEventListener('click', () => {
      run(() => Promise.resolve(), 'validacao atualizada');
    });

    updateClock();
    window.setInterval(updateClock, 1000);
    load().catch((error) => setStatus(error.message, 'bad'));
  </script>
</body>
</html>
`
