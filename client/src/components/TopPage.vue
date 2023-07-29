<template>
  <div>
    <h1>Leaning Othello</h1>
    <div><router-link to="/playBoard">PlayOthello</router-link></div>
    <div><router-link to="/setting">setting</router-link></div>
  </div>
</template>

<script>
export default {
  name: 'SelectMenu',
  data: function () {
    return {
      info: null,
    }
  },
  methods: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
