import Vue from 'vue'
import Router from 'vue-router'
import axios from 'axios'
import literalReaction from '@/components/literalReaction'
import result from '@/components/result'
import favs from '@/components/favs'

Vue.use(Router)

Vue.prototype.$axios = axios

export default new Router({
  routes: [
    {
      path: '/',
      name: 'literalReaction',
      component: literalReaction
    },
    {
      path: '/result/adv-v',
      name: 'result',
      component: result,
      props: {
        formerPOS: 'adverb',
        latterPOS: 'verb'
      }
    },
    {
      path: '/result/adj-n',
      name: 'result',
      component: result,
      props: {
        formerPOS: 'adj',
        latterPOS: 'noun'
      }
    },
    {
      path: '/result/n-v',
      name: 'result',
      component: result,
      props: {
        formerPOS: 'noun',
        latterPOS: 'verb'
      }
    },
    {
      path: '/favs',
      name: 'favs',
      component: favs
    }
  ]
})
