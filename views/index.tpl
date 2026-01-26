<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Property Listing API</title>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Fraunces:opsz,wght@9..144,600;9..144,700&family=IBM+Plex+Mono:wght@400;500&family=Work+Sans:wght@400;600&display=swap" rel="stylesheet">
  <style>
    :root {
      --bg: #f7f3ea;
      --ink: #1b1a17;
      --muted: #5e5b55;
      --accent: #c65b2e;
      --accent-2: #2a6a7a;
      --paper: #fff9ef;
      --shadow: 0 24px 70px rgba(27, 26, 23, 0.18);
    }

    * {
      box-sizing: border-box;
    }

    body {
      margin: 0;
      font-family: "Work Sans", "Trebuchet MS", sans-serif;
      color: var(--ink);
      background: radial-gradient(circle at 15% 20%, #ffe7cf 0%, transparent 45%),
        radial-gradient(circle at 80% 10%, #dff1f6 0%, transparent 40%),
        linear-gradient(160deg, #f8f0e2 0%, #fff7ec 40%, #f6efe2 100%);
      min-height: 100vh;
    }

    .page {
      max-width: 1100px;
      margin: 0 auto;
      padding: 64px 24px 72px;
    }

    header {
      display: grid;
      gap: 24px;
      align-items: center;
    }

    .badge {
      display: inline-flex;
      align-items: center;
      gap: 10px;
      font-size: 13px;
      letter-spacing: 0.2em;
      text-transform: uppercase;
      color: var(--accent-2);
      font-weight: 600;
    }

    .badge::before {
      content: "";
      width: 36px;
      height: 2px;
      background: var(--accent-2);
    }

    h1 {
      font-family: "Fraunces", "Georgia", serif;
      font-size: clamp(2.4rem, 4vw, 4.2rem);
      margin: 0;
    }

    .lead {
      font-size: 1.1rem;
      color: var(--muted);
      max-width: 680px;
      line-height: 1.7;
    }

    .grid {
      margin-top: 48px;
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
      gap: 24px;
    }

    .card {
      background: var(--paper);
      padding: 22px 22px 24px;
      border-radius: 16px;
      border: 1px solid rgba(27, 26, 23, 0.08);
      box-shadow: var(--shadow);
      min-height: 150px;
      position: relative;
      overflow: hidden;
    }

    .card::after {
      content: "";
      position: absolute;
      inset: auto -30% -60% -30%;
      height: 100px;
      background: linear-gradient(120deg, rgba(198, 91, 46, 0.15), transparent);
      transform: rotate(-2deg);
    }

    .card h3 {
      margin: 0 0 10px;
      font-size: 1.05rem;
      font-weight: 600;
    }

    .card p {
      margin: 0;
      color: var(--muted);
      line-height: 1.6;
    }

    .code {
      margin-top: 32px;
      background: #1d1b18;
      color: #f4efe7;
      padding: 18px 20px;
      border-radius: 14px;
      font-family: "IBM Plex Mono", ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
      font-size: 0.92rem;
      box-shadow: var(--shadow);
      overflow-x: auto;
    }

    .meta {
      display: flex;
      flex-wrap: wrap;
      gap: 12px 24px;
      margin-top: 28px;
      font-size: 0.95rem;
    }

    .meta span {
      color: var(--muted);
    }

    footer {
      margin-top: 52px;
      display: flex;
      flex-wrap: wrap;
      gap: 12px 28px;
      font-size: 0.92rem;
      color: var(--muted);
    }

    footer a {
      color: var(--accent);
      text-decoration: none;
      font-weight: 600;
    }

    @media (max-width: 640px) {
      .page {
        padding: 48px 18px 60px;
      }

      .card {
        box-shadow: none;
      }

      header {
        gap: 16px;
      }
    }
  </style>
</head>
<body>
  <div class="page">
    <header>
      <div class="badge">Property Listing API</div>
      <h1>Welcome. Discover property data with one clean endpoint.</h1>
      
    </header>

    <section class="grid">
      <div class="card">
        <h3>Location-driven results</h3>
        <p>Query by location and receive curated property details in a consistent format.</p>
      </div>
      
      <div class="card">
        <h3>Secure by API key</h3>
        <p>Authenticate via the <code>x-api-key</code> header to access the endpoint.</p>
      </div>
    </section>

    <div class="code">
      GET /api/properties?location=usa:florida:destin&amp;items=true
    </div>

    <div class="meta">
      <span><strong>Headers:</strong> x-api-key: &lt;YOUR_API_KEY&gt;</span>
      <span><strong>Response:</strong> JSON list of normalized properties</span>
    </div>

    <footer>
      <div>
        <h5>Copyrights © 2026 All Rights Reserved </h5>
      </div>
      <div></div>
    </footer>
  </div>
</body>
</html>
