<template>
  <div id="app">
    <router-link to="/" tag="header" id="header">
      <h1 id="headertitle"><vue-fontawesome icon="flask"></vue-fontawesome> Literal Reaction</h1>
    </router-link>
    <!-- show each page component here -->
    <router-view><router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  width: 100%;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #330055;
}

#header {
  text-align: left;
  width: 100%;
  /*    height width */
  padding-top: 0.1rem;
  padding-bottom: 0.3rem;
  margin: 0 auto;
  background-color: #a67dff;
}

#headertitle {
  color: #ffffff;
  weight: bold;
  margin-left: 2rem;
  margin-top: 0.1rem;
  margin-bottom: 0.1rem;
  font-size: 1.5rem;
}

.left {
  text-align: left;
  margin-left: 1rem;
}

.btn {
  display: block;
  padding: 0.3rem 1rem;
  margin: 0.5rem auto;
  width: 35%;
  text-decoration: none;
  font-size: 1.5rem;
  background: #ffffff;
  color: #222222;
  border: solid 2px #222222;
  border-radius: 4px;
  transition: .25s;
}

.btn:hover {
  background: #222222;
  color: #ffffff;
}

.back-btn {
  width: 15%
}
</style>
