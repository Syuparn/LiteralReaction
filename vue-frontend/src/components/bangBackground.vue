<template>
  <div v-bind:style="bangStyle">
    <slot name="inner"></slot>
  </div>
</template>

<script>
export default {
  name: 'bangBackground',
  data: function () {
    return {
      bangSVG: new window.ConicGradient({
        repeating: true,
        stops: `#ffffff 0,
                #ffffff 2.0%,
                #AAAAAA 2.125%,
                #AAAAAA 2.375%,
                #ffffff 2.5%`
      })
    }
  },
  computed: {
    bangStyle: function () {
      return {
        'background-image': `url("${this.bangSVG.blobURL}")`
      }
    }
  }
}
</script>
