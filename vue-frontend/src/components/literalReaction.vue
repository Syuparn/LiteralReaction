<template>
  <div>
    <h1 id="title"><vue-fontawesome icon="flask"></vue-fontawesome> Literal Reaction</h1>
    <h2 id="subtitle">〜言語反応〜</h2>
    <div id="btn-list">
      <router-link to="/result/adj-n" class="btn">Adjective + Noun → Phrase</router-link>
      <router-link to="/result/adv-v" class="btn">Adverb + Verb → Phrase</router-link>
      <router-link to="/result/n-v" class="btn">Noun + Particle + Verb → Phrase</router-link>
      <router-link to="/favs" class="btn">Favorites</router-link>
    </div>
  </div>
</template>

<script>
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#title {
  font-weight: bold;
  font-size: 3rem;
  margin-bottom: 0rem;
}

#subtitle {
  font-size: 1.5rem;
  font-weight: bold;
}
</style>
