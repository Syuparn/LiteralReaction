<template>
  <div>
    <h1>みんなのお気に入り</h1>
    <h2 v-if="err">通信エラー…もう一度試してみてね</h2>
    <ul v-else>
      <li v-for="phrase in phrases" v-bind:key="phrase.id">{{phrase.str}}</li>
    </ul>
    <div v-if="!fetched">Now loading...</div>
    <button class="btn back-btn" v-on:click="more">
      More
    </button>
    <router-link to="/" class="btn back-btn">Back</router-link>
  </div>
</template>

<script>
const PAGE_SIZE = 100

export default {
  name: 'favs',
  props: [],
  data: function () {
    return {
      err: false,
      fetched: false,
      page: 1,
      unfetchedDataExists: true,
      phrases: []
    }
  },
  created: function () {
    this.more()
  },
  methods: {
    more: function () {
      this.$axios.get(`/api/favs/${this.page}`)
        .then((response) => {
          console.log('status:', response.status)
          console.log('body:', response.data)
          var phrases = response.data.map((d, i) => ({
            id: i,
            str: d.FormerWord + d.Particle + d.LatterWord
          }))

          if (phrases.length < PAGE_SIZE) {
            this.unfetchedDataExists = false
          }

          this.phrases = this.phrases.concat(phrases)
          this.fetched = true
          this.page += 1
        })
        .catch((reason) => {
          this.err = true
          console.log('failed to post fav:', reason)
        })
    }
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
  /* align li horizontally until the row gets full */
  flex-direction:row;
  flex-wrap:wrap;
  justify-content: start;
  /* include border width in box size */
  box-sizing: border-box;
  width:100%;
}

li {
  /* 2 columns in each line */
  width:40%;
  /* include border width in box size */
  box-sizing: border-box;
  padding-top: 1rem;
  padding-bottom: 1rem;
  margin-left: 5%;
  margin-right: 5%;
  margin-top: 2.5%;
  margin-bottom: 2.5%;
  background: #a67dff;
  color: #ffffff;
  font-size: 2.5rem;
  weight: bold;
}
</style>
