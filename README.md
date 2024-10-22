<h1 align="center">Go Integrator</h1>
<p align="center">
  Example of integrating different applications for data exchange.
  <br/>
  <strong>(work in progress)</strong>
  <br/><br/>
  <a href="https://github.com/betofigueiredo/go-integrator/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge&labelColor=363a4f&color=a6da95"></a>
</p>

<h2>Steps</h2>

<h3>🔹Request users list from API</h3>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Concurrently, in chunks of 1000
<br/>
<h3>🔹Map all users ID’s</h3>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Save in a map, using mutex to handle async r/w
<br/>
<h3>🔹Request additional info for each user</h3>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Concurrently, in chunks to prevent API overflow
<br/>
<h3>🔹Save additional info in the map</h3>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Also using mutex to handle async r/w
<br/>
<h3>🔹Send users data</h3>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;After processed, send to any selected external service
<br/><br/><br/>

<h2>API</h2>

<p>
  <a href="https://www.python.org/"><img src="https://img.shields.io/badge/Made%20with-Python-blue?style=for-the-badge&labelColor=363a4f&color=346FA0"></a>
  <a href="https://fastapi.tiangolo.com/"><img src="https://img.shields.io/badge/Made%20with-FastAPI-blue?style=for-the-badge&labelColor=363a4f&color=009485"></a>
  <br/><br/>
  Endpoint: <strong>/users</strong>
  <br/>
  <img src="https://github.com/user-attachments/assets/e4adfcf2-29db-4cb0-99ad-1952c7e9708c" alt="API Schema 1" />
  <br/><br/>
  Endpoint: <strong>/users/{user_id}</strong>
  <br/>
  <img src="https://github.com/user-attachments/assets/5b2f2487-422c-408c-aabf-2bbd5028427d" alt="API Schema 2" />
  <br/><br/><br/>
</p>

<h2>Integrator</h2>

<p>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Made%20with-Go-blue?style=for-the-badge&labelColor=363a4f&color=007d9c"></a>
  <a href="https://gofiber.io/"><img src="https://img.shields.io/badge/Made%20with-Fiber-blue?style=for-the-badge&labelColor=363a4f&color=1F4F98"></a>
  <br/><br/>
  Endpoint: <strong>/get-users</strong>
  <br/>
  <img src="https://github.com/user-attachments/assets/f7c1ea85-2866-48c7-989f-958bebb684e5" alt="INTEGRATOR Schema" />
  <br/><br/><br/>
</p>

<h2> :zap: Usage</h2>

```zsh
❯ gh repo clone ...

❯ make up

...
```
