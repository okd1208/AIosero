<template>
  <div class="hello">
    <h1>Leaning Osero</h1>
    <button @click="Send">PlayOsero</button>
    <router-link to="/playBoard">play</router-link>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'HelloWorld',
  data: function () {
    return {
      info: null,
    }
  },
  methods: {
    async Send() {
      axios
      .get('http://localhost:8888')
      .then(response => (this.info = response))
      console.log(this.info);
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
