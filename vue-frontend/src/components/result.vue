<template>
  <div>
    <div v-if="err">
      <h1>反応失敗…</h1>
      <p>（通信エラー） もう一度試してみてね</p>
    </div>
    <bang-background v-else>
      <template v-slot:inner>
        <h1 class="left">結果：</h1>
        <table>
          <tbody>
            <tr>
              <th id="former">{{former}}</th>
              <th v-if="requiresParticle" id="particle">
                <a href="#" v-on:click="changeParticle" id="particle-btn"
                 v-bind:class="[particleClass]">{{particle}}</a>
              </th>
              <th id="latter">{{latter}}</th>
            </tr>
          </tbody>
        </table>
        <button class="btn fav-btn" v-on:click="fav"
         v-bind:class="{'fav-btn-clicked': isFav}">
          <vue-fontawesome icon="thumbs-up"></vue-fontawesome>
        </button>
      </template>
    </bang-background>
    <router-link to="/" class="btn back-btn">Back</router-link>
  </div>
</template>

<script>
import bangBackground from './bangBackground'

export default {
  name: 'result',
  components: {
    'bang-background': bangBackground
  },
  props: ['formerPOS', 'latterPOS'],
  data: function () {
    return {
      former: '？？？',
      latter: '？？？',
      particleID: 0,
      particles: ['を', 'に', 'が', 'で', 'と'],
      hoversOnParticle: false,
      isFav: false,
      err: false
    }
  },
  computed: {
    requiresParticle: function () {
      return this.formerPOS === 'noun' && this.latterPOS === 'verb'
    },
    particle: function () {
      return this.particles[this.particleID]
    },
    particleClass: function () {
      return [
        'green-btn',
        'yellow-btn',
        'purple-btn',
        'orange-btn',
        'gray-btn'
      ][this.particleID]
    }
  },
  created: function () {
    this.$axios.get(`/api/rand/${this.formerPOS}`)
      .then((response) => {
        console.log('status:', response.status)
        console.log('body:', response.data)
        this.former = response.data.Word
      })
      .catch((reason) => {
        this.err = true
      })

    this.$axios.get(`/api/rand/${this.latterPOS}`)
      .then((response) => {
        console.log('status:', response.status)
        console.log('body:', response.data)
        this.latter = response.data.Word
      })
      .catch((reason) => {
        this.err = true
      })
  },
  methods: {
    changeParticle: function () {
      this.particleID = (this.particleID + 1) % this.particles.length
    },
    fav: function () {
      // TODO: api post if !this.isFav
      this.isFav = true
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#former {
  color: #dc143c;
  font-size:4rem;
}

#latter {
  color: #0000cd;
  font-size:4rem;
}

table {
  /*    height width */
  margin: 0.3rem auto;
  border-collapse: separate;
  border-spacing: 3rem;
}

#particle-btn {
  width: 2.5rem;
  display: inline-block;
  padding: 1rem 1rem;
  margin: 0.5rem auto;
  text-decoration: none;
  font-size: 2rem;
  border-radius: 4px;
}

.fav-btn {
  width: 25%;
  padding: 0.2rem 0.2rem;
  color: #a67dff;
  font-size: 2.5rem;
  border: solid 2px #a67dff;
}

.fav-btn:hover, .fav-btn-clicked {
  background: #a67dff;
  color: #ffffff;
}

.green-btn {
  background: #ffffff;
  color: #228b22;
  border: solid 2px #228b22;
}

.green-btn:hover {
  background: #228b22;
  color: #ffffff;
}

.yellow-btn {
  background: #ffffff;
  color: #ffa500;
  border: solid 2px #ffa500;
}

.yellow-btn:hover {
  background: #ffa500;
  color: #ffffff;
}

.purple-btn {
  background: #ffffff;
  color: #4b0082;
  border: solid 2px #4b0082;
}

.purple-btn:hover {
  background: #4b0082;
  color: #ffffff;
}

.orange-btn {
  background: #ffffff;
  color: #ff7343;
  border: solid 2px #ff7343;
}

.orange-btn:hover {
  background: #ff7343;
  color: #ffffff;
}

.gray-btn {
  background: #ffffff;
  color: #444444;
  border: solid 2px #444444;
}

.gray-btn:hover {
  background: #444444;
  color: #ffffff;
}
</style>
