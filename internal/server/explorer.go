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

    .visual-panel {
      padding: 0;
      overflow: hidden;
    }

    .visual-head {
      display: flex;
      justify-content: space-between;
      gap: 14px;
      align-items: center;
      padding: 18px 18px 0;
    }

    .visual-head p {
      color: var(--muted);
      font-size: 13px;
      margin-top: 4px;
    }

    .visual-badge {
      border: 1px solid var(--line);
      border-radius: 999px;
      color: var(--muted);
      padding: 6px 10px;
      font-size: 12px;
      white-space: nowrap;
    }

    .chain-canvas-wrap {
      position: relative;
      height: 330px;
      margin-top: 10px;
      border-top: 1px solid rgba(36, 53, 83, 0.62);
      background:
        linear-gradient(rgba(229, 237, 248, 0.035) 1px, transparent 1px),
        linear-gradient(90deg, rgba(229, 237, 248, 0.035) 1px, transparent 1px),
        radial-gradient(circle at 50% 30%, rgba(0, 173, 216, 0.12), transparent 26rem);
      background-size: 34px 34px, 34px 34px, auto;
    }

    #chain-canvas {
      width: 100%;
      height: 100%;
      display: block;
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

      .visual-head {
        align-items: start;
        display: grid;
      }

      .chain-canvas-wrap {
        height: 290px;
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

    <section class="panel visual-panel">
      <div class="visual-head">
        <div>
          <h2>Chain Visualizer</h2>
          <p>Blocks are cubes; hashes are the chain links between them.</p>
        </div>
        <span class="visual-badge" id="visual-status">waiting for chain</span>
      </div>
      <div class="chain-canvas-wrap">
        <canvas id="chain-canvas" width="1120" height="330" aria-label="Animated blockchain cubes"></canvas>
      </div>
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
      busy: false,
      chain: [],
      validation: { valid: true },
      animationStarted: false,
      visualTime: 0
    };

    const $ = (id) => document.getElementById(id);
    const shortHash = (hash) => hash ? hash.slice(0, 14) + '...' + hash.slice(-8) : '-';
    const formatDate = (value) => value ? new Date(value).toLocaleString() : '-';
    const canvas = $('chain-canvas');
    const ctx = canvas.getContext('2d');

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

    function resizeCanvas() {
      const rect = canvas.getBoundingClientRect();
      const ratio = window.devicePixelRatio || 1;
      canvas.width = Math.max(1, Math.floor(rect.width * ratio));
      canvas.height = Math.max(1, Math.floor(rect.height * ratio));
      ctx.setTransform(ratio, 0, 0, ratio, 0, 0);
    }

    function cubePoints(x, y, width, height, depth) {
      return {
        top: [
          [x, y - height],
          [x + width * 0.5, y - height - depth],
          [x + width, y - height],
          [x + width * 0.5, y - height + depth]
        ],
        left: [
          [x, y - height],
          [x + width * 0.5, y - height + depth],
          [x + width * 0.5, y + depth],
          [x, y]
        ],
        right: [
          [x + width, y - height],
          [x + width * 0.5, y - height + depth],
          [x + width * 0.5, y + depth],
          [x + width, y]
        ]
      };
    }

    function drawPolygon(points, fill, stroke = 'rgba(229, 237, 248, 0.18)') {
      ctx.beginPath();
      points.forEach(([x, y], index) => {
        if (index === 0) ctx.moveTo(x, y);
        else ctx.lineTo(x, y);
      });
      ctx.closePath();
      ctx.fillStyle = fill;
      ctx.fill();
      ctx.strokeStyle = stroke;
      ctx.lineWidth = 1;
      ctx.stroke();
    }

    function drawCube(block, x, y, size, valid, phase) {
      const depth = size * 0.28;
      const bob = Math.sin(phase + block.index * 0.8) * 4;
      const points = cubePoints(x, y + bob, size, size * 0.72, depth);
      const accent = valid ? '0, 173, 216' : '249, 112, 102';

      ctx.shadowColor = 'rgba(' + accent + ', 0.36)';
      ctx.shadowBlur = 18;
      drawPolygon(points.top, 'rgba(' + accent + ', 0.42)');
      ctx.shadowBlur = 0;
      drawPolygon(points.left, 'rgba(16, 26, 44, 0.98)');
      drawPolygon(points.right, 'rgba(' + accent + ', 0.22)');

      ctx.fillStyle = valid ? '#e5edf8' : '#ffe2df';
      ctx.font = '700 14px ui-monospace, monospace';
      ctx.textAlign = 'center';
      ctx.fillText('#' + block.index, x + size * 0.5, y + bob - size * 0.28);

      ctx.fillStyle = valid ? '#8ba0ba' : '#f97066';
      ctx.font = '11px ui-monospace, monospace';
      ctx.fillText(shortHash(block.hash), x + size * 0.5, y + bob + depth + 18);
    }

    function drawChainLink(x1, y1, x2, y2, index, valid, phase) {
      const dx = x2 - x1;
      const dy = y2 - y1;
      const angle = Math.atan2(dy, dx);
      const distance = Math.hypot(dx, dy);
      const links = Math.max(2, Math.floor(distance / 32));
      const accent = valid ? '#00add8' : '#f97066';

      ctx.save();
      ctx.translate(x1, y1);
      ctx.rotate(angle);

      for (let i = 0; i < links; i++) {
        const progress = (i + 0.5) / links;
        const sway = Math.sin(phase * 1.8 + i * 0.7 + index) * 5;
        const x = distance * progress;

        ctx.save();
        ctx.translate(x, sway);
        ctx.rotate(i % 2 === 0 ? 0.35 : -0.35);
        ctx.beginPath();
        ctx.ellipse(0, 0, 13, 7, 0, 0, Math.PI * 2);
        ctx.strokeStyle = accent;
        ctx.globalAlpha = valid ? 0.78 : 0.95;
        ctx.lineWidth = 3;
        ctx.stroke();
        ctx.restore();
      }

      ctx.restore();
    }

    function visibleBlocks(blocks) {
      if (blocks.length <= 7) return blocks;
      return blocks.slice(-7);
    }

    function drawVisualizer() {
      resizeCanvas();
      const rect = canvas.getBoundingClientRect();
      const width = rect.width;
      const height = rect.height;
      const blocks = visibleBlocks(state.chain);
      const valid = state.validation.valid !== false;
      const phase = state.visualTime;

      ctx.clearRect(0, 0, width, height);

      ctx.fillStyle = 'rgba(5, 12, 24, 0.34)';
      ctx.fillRect(0, 0, width, height);

      ctx.strokeStyle = 'rgba(139, 160, 186, 0.18)';
      ctx.lineWidth = 1;
      ctx.beginPath();
      ctx.moveTo(24, height * 0.72);
      ctx.bezierCurveTo(width * 0.28, height * 0.58, width * 0.62, height * 0.86, width - 24, height * 0.68);
      ctx.stroke();

      if (!blocks.length) {
        ctx.fillStyle = '#8ba0ba';
        ctx.font = '14px ui-monospace, monospace';
        ctx.textAlign = 'center';
        ctx.fillText('no blocks yet', width * 0.5, height * 0.5);
        return;
      }

      const gap = width / (blocks.length + 0.8);
      const size = Math.max(54, Math.min(92, gap * 0.58));
      const positions = blocks.map((block, index) => ({
        block,
        x: gap * (index + 0.42),
        y: height * 0.64 + Math.sin(index * 0.9) * 12,
        size
      }));

      positions.forEach((current, index) => {
        const next = positions[index + 1];
        if (!next) return;

        drawChainLink(
          current.x + current.size * 0.92,
          current.y - current.size * 0.26,
          next.x + next.size * 0.08,
          next.y - next.size * 0.24,
          index,
          valid,
          phase
        );
      });

      positions.forEach(({ block, x, y, size }) => {
        drawCube(block, x, y, size, valid, phase);
      });

      if (state.chain.length > blocks.length) {
        ctx.fillStyle = '#8ba0ba';
        ctx.font = '12px ui-monospace, monospace';
        ctx.textAlign = 'left';
        ctx.fillText('showing latest ' + blocks.length + ' of ' + state.chain.length + ' blocks', 18, height - 18);
      }
    }

    function animateVisualizer() {
      state.visualTime += 0.016;
      drawVisualizer();
      window.requestAnimationFrame(animateVisualizer);
    }

    function startVisualizer() {
      if (state.animationStarted) return;
      state.animationStarted = true;
      resizeCanvas();
      window.addEventListener('resize', resizeCanvas);
      window.requestAnimationFrame(animateVisualizer);
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
      $('visual-status').textContent = validation.valid ? 'chain links locked' : 'broken chain detected';
      $('visual-status').style.color = validation.valid ? 'var(--muted)' : 'var(--bad)';

      state.chain = chain;
      state.validation = validation;
      renderBlocks(chain);
      renderPending(pending);
      startVisualizer();
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
