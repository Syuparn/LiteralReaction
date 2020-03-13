<template>
  <div>
    <h1>みんなのお気に入り</h1>
    <h2 v-if="err">通信エラー…もう一度試してみてね</h2>
    <ul v-else>
      <li v-for="phrase in phrases" v-bind:key="phrase.id">{{phrase.str}}</li>
    </ul>
    <div v-if="!fetched">Now loading...</div>
    <router-link to="/" class="btn back-btn">Back</router-link>
  </div>
</template>

<script>
export default {
  name: 'favs',
  props: [],
  data: function () {
    return {
      err: false,
      fetched: false,
      phrases: []
    }
  },
  created: function () {
    this.phrases = [{id: 0, str: '深い謎'}, {id: 1, str: '淡い期待'}]
    this.fetched = true
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style: none;
  display: flex;
}

li {
  padding: 1rem;
  margin: 1rem;
  background: #a67dff;
  color: #ffffff;
  font-size: 2.5rem;
  weight: bold;
}
</style>
