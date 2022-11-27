<template>
  <div class="hello">
    <h1>Leaning Osero testt</h1>
    <button @click="Send">PlayOsero</button>
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
