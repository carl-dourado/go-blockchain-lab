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
      --panel: #111616;
      --panel-strong: #202020;
      --text: #f4f4f4;
      --muted: #a7b4b4;
      --line: #77d7e8;
      --accent: #1499ff;
      --blue: #1499ff;
      --yellow: #ffff00;
      --red: #ff2020;
      --ok: #65ff8f;
      --warn: #ffff00;
      --bad: #ff4d4d;
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace;
    }

    * {
      box-sizing: border-box;
    }

    body {
      margin: 0;
      min-height: 100vh;
      background:
        radial-gradient(circle at left top, rgba(20, 153, 255, 0.24), transparent 23rem),
        radial-gradient(circle at right bottom, rgba(119, 215, 232, 0.12), transparent 28rem),
        var(--bg);
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
      margin: 0 auto;
      min-height: calc(100vh - 10px);
      padding: 24px 24px 18px;
      background: var(--screen);
      border: 1px solid rgba(119, 215, 232, 0.92);
      box-shadow:
        0 0 0 1px rgba(20, 153, 255, 0.14),
        0 0 30px rgba(20, 153, 255, 0.22),
        inset 0 0 32px rgba(20, 153, 255, 0.04);
      border-radius: 10px 10px 0 0;
    }

    header {
      margin-bottom: 20px;
    }

    h1,
    h2,
    h3,
    p {
      margin: 0;
    }

    .eyebrow {
      color: var(--text);
      font-size: 16px;
      letter-spacing: 0;
    }

    h1 {
      margin-top: 8px;
      font-size: clamp(24px, 5vw, 50px);
      line-height: 0.95;
    }

    .hero-copy {
      color: var(--yellow);
      margin-top: 10px;
      max-width: 760px;
      line-height: 1.55;
    }

    .terminal-menu {
      display: flex;
      justify-content: space-between;
      gap: 16px;
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

    .terminal-clock {
      min-width: 170px;
      justify-items: start;
    }

    .terminal-clock span {
      display: grid;
      grid-template-columns: 50px 1fr;
      gap: 10px;
    }

    .ascii-stage {
      display: grid;
      grid-template-columns: minmax(180px, 0.52fr) minmax(320px, 1fr);
      gap: clamp(20px, 5vw, 54px);
      align-items: end;
      margin-bottom: 18px;
    }

    .ascii-cat,
    .ascii-logo {
      margin: 0;
      white-space: pre;
      line-height: 1.12;
      font-size: clamp(10px, 1.55vw, 18px);
    }

    .ascii-cat {
      color: var(--text);
    }

    .ascii-logo {
      font-weight: 800;
      font-size: clamp(12px, 1.9vw, 24px);
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
      margin-bottom: 18px;
      line-height: 1.35;
    }

    .terminal-subtitle h1 {
      color: var(--yellow);
      font-size: clamp(18px, 2.5vw, 24px);
      margin: 0;
    }

    .terminal-yellow {
      color: var(--yellow);
      font-weight: 800;
    }

    .terminal-blue {
      color: var(--blue);
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

    .action-grid {
      display: grid;
      grid-template-columns: 1.2fr 0.8fr 1fr;
      gap: 14px;
      margin-bottom: 14px;
    }

    .panel,
    .stat,
    .action-card,
    .block,
    .record {
      border: 1px solid var(--line);
      background: rgba(8, 12, 12, 0.86);
      border-radius: 0;
    }

    .stat {
      min-height: 92px;
      padding: 14px;
    }

    .stat span {
      display: block;
      color: var(--blue);
      font-size: 12px;
      margin-bottom: 10px;
    }

    .stat strong {
      font-size: 26px;
      color: var(--yellow);
      overflow-wrap: anywhere;
    }

    .action-card {
      padding: 16px;
    }

    .step-label {
      display: inline-flex;
      align-items: center;
      min-height: 26px;
      border: 1px solid rgba(20, 153, 255, 0.72);
      border-radius: 0;
      color: var(--yellow);
      padding: 4px 9px;
      font-size: 12px;
      margin-bottom: 12px;
    }

    .action-card h2 {
      font-size: 18px;
      margin-bottom: 6px;
      color: var(--blue);
    }

    .action-card p {
      color: var(--text);
      font-size: 13px;
      line-height: 1.45;
      margin-bottom: 14px;
    }

    .action-card strong {
      color: var(--text);
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

    .visual-head h2 {
      font-size: 22px;
      color: var(--blue);
    }

    .visual-head p {
      color: var(--yellow);
      font-size: 13px;
      margin-top: 4px;
      line-height: 1.45;
    }

    .visual-badge {
      border: 1px solid var(--line);
      border-radius: 0;
      color: var(--yellow);
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
        linear-gradient(rgba(20, 153, 255, 0.05) 1px, transparent 1px),
        linear-gradient(90deg, rgba(20, 153, 255, 0.05) 1px, transparent 1px),
        radial-gradient(circle at 50% 30%, rgba(20, 153, 255, 0.12), transparent 24rem),
        #101010;
      background-size: 28px 28px, 28px 28px, auto, auto;
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
      color: var(--yellow);
      font-size: 13px;
      margin-top: 4px;
      line-height: 1.45;
    }

    form {
      display: grid;
      grid-template-columns: 1fr;
      gap: 10px;
    }

    label {
      display: grid;
      gap: 6px;
      color: var(--blue);
      font-size: 12px;
    }

    input,
    textarea {
      width: 100%;
      border: 1px solid var(--line);
      border-radius: 0;
      background: #050505;
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
      border-radius: 0;
      color: #050505;
      background: var(--yellow);
      padding: 12px 14px;
      cursor: pointer;
      min-height: 44px;
      width: fit-content;
    }

    button.secondary {
      background: transparent;
      color: var(--yellow);
      border-color: var(--line);
    }

    button:disabled {
      opacity: 0.55;
      cursor: wait;
    }

    .status {
      min-height: 24px;
      color: var(--yellow);
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
      border-radius: 0;
      background: var(--panel-strong);
      color: var(--yellow);
      font-weight: 800;
    }

    .hash,
    .muted {
      color: var(--muted);
      overflow-wrap: anywhere;
    }

    .hash {
      font-size: 12px;
      color: var(--blue);
    }

    .pill {
      display: inline-flex;
      align-items: center;
      border-radius: 0;
      border: 1px solid var(--line);
      color: var(--yellow);
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
      background: #0b0b0b;
    }

    .record b {
      color: var(--text);
      overflow-wrap: anywhere;
    }

    .empty {
      color: var(--muted);
      border: 1px dashed var(--line);
      border-radius: 0;
      padding: 18px;
      text-align: center;
    }

    .terminal-status {
      display: flex;
      justify-content: space-between;
      gap: 18px;
      margin-top: 18px;
      color: var(--text);
      font-size: 15px;
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
      header,
      .layout {
        display: grid;
      }

      .toolbar {
        justify-content: start;
      }

      .terminal-menu,
      .ascii-stage,
      .terminal-status {
        display: grid;
      }

      .terminal-clock {
        justify-items: start;
      }

      .ascii-logo {
        font-size: clamp(10px, 4vw, 18px);
        overflow-x: auto;
      }

      .visual-head {
        align-items: start;
        display: grid;
      }

      .chain-canvas-wrap {
        height: 290px;
      }

      .grid,
      .action-grid,
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
          <button class="secondary" id="refresh" type="button">refresh</button>
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
        <h1>Go Blockchain Lab</h1>
        <span class="terminal-yellow">The GOCHAIN 0.1 local blockchain system</span>
        <span class="terminal-blue">records -> pending queue -> mine block -> validate previous hash</span>
      </div>
    </header>

    <section class="panel visual-panel">
      <div class="visual-head">
        <div>
          <h2>Corrente de blocos</h2>
          <p>Cada cubo e um bloco. Os elos representam a ligacao pelo hash anterior.</p>
        </div>
        <span class="visual-badge" id="visual-status">carregando cadeia</span>
      </div>
      <div class="chain-canvas-wrap">
        <canvas id="chain-canvas" width="1120" height="330" aria-label="Animated blockchain cubes"></canvas>
      </div>
    </section>

    <section class="grid" aria-label="Resumo da cadeia">
      <article class="stat"><span>altura</span><strong id="height">-</strong></article>
      <article class="stat"><span>pendentes</span><strong id="pending-count">-</strong></article>
      <article class="stat"><span>validacao</span><strong id="valid">-</strong></article>
      <article class="stat"><span>ultimo hash</span><strong class="hash" id="latest-hash">-</strong></article>
    </section>

    <section class="action-grid" aria-label="Fluxo do lab">
      <article class="action-card">
        <span class="step-label">1. registro</span>
        <h2>Criar dado pendente</h2>
        <p>O dado entra na fila. Ele ainda nao faz parte da cadeia.</p>
        <form id="record-form">
          <label>
            autor
            <input id="author" value="carlos" autocomplete="off">
          </label>
          <label>
            dado
            <textarea id="data">transfer 10 from alice to bob</textarea>
          </label>
          <button type="submit">adicionar</button>
        </form>
        <p class="status" id="status"></p>
      </article>

      <article class="action-card">
        <span class="step-label">2. mineracao</span>
        <h2>Fechar novo bloco</h2>
        <p><strong id="mine-count">0</strong> registros esperam para virar bloco.</p>
        <button id="mine" type="button">minerar pendentes</button>
      </article>

      <article class="action-card">
        <span class="step-label">3. validacao</span>
        <h2>Conferir corrente</h2>
        <p id="validation-note">A cadeia sera recalculada pelos hashes.</p>
        <button class="secondary" id="validate-now" type="button">validar agora</button>
      </article>
    </section>

    <section class="layout">
      <section class="panel">
        <div class="panel-head">
          <div>
            <h2>Blocos minerados</h2>
            <p>O bloco mais novo aparece primeiro. O hash anterior aponta para o bloco de baixo.</p>
          </div>
        </div>
        <div class="list" id="blocks"></div>
      </section>

      <aside class="panel">
        <div class="panel-head">
          <div>
            <h2>Fila pendente</h2>
            <p>Estes registros ainda nao foram minerados.</p>
          </div>
        </div>
        <div class="list" id="pending"></div>
      </aside>
    </section>

    <footer class="terminal-status">
      <span>Logon ===&gt; <b id="prompt-status">READY</b><i class="prompt-cursor" aria-hidden="true"></i></span>
      <span>RUNNING&nbsp;&nbsp;GOCHAIN</span>
    </footer>
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

    function pad(value) {
      return String(value).padStart(2, '0');
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
        empty.textContent = 'nenhum bloco encontrado';
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
        title.textContent = block.index === 0 ? 'bloco genesis' : 'bloco #' + block.index;
        text.appendChild(title);

        const hash = document.createElement('p');
        hash.className = 'hash';
        hash.textContent = 'hash atual: ' + block.hash;
        text.appendChild(hash);
        top.appendChild(text);

        const nonce = document.createElement('span');
        nonce.className = 'pill';
        nonce.textContent = 'nonce ' + block.nonce;
        top.appendChild(nonce);

        item.appendChild(top);

        const prev = document.createElement('p');
        prev.className = 'hash';
        prev.textContent = 'hash anterior: ' + block.previousHash;
        item.appendChild(prev);

        const meta = document.createElement('p');
        meta.className = 'muted';
        meta.textContent = formatDate(block.timestamp) + ' | dificuldade ' + block.difficulty + ' | registros ' + block.records.length;
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
        empty.textContent = 'sem registros pendentes';
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
        ctx.fillText('nenhum bloco ainda', width * 0.5, height * 0.5);
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
        ctx.fillText('mostrando os ultimos ' + blocks.length + ' de ' + state.chain.length + ' blocos', 18, height - 18);
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
      $('mine-count').textContent = pending.length;
      $('valid').textContent = validation.valid ? 'ok' : 'falhou';
      $('valid').style.color = validation.valid ? 'var(--ok)' : 'var(--bad)';
      $('latest-hash').textContent = shortHash(latest && latest.hash);
      $('visual-status').textContent = validation.valid ? 'corrente valida' : 'corrente quebrada';
      $('visual-status').style.color = validation.valid ? 'var(--muted)' : 'var(--bad)';
      $('validation-note').textContent = validation.valid
        ? 'Hashes e ligacoes batem. Nenhum bloco parece adulterado.'
        : 'Algum hash ou elo nao bate. A cadeia precisa ser revisada.';

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
