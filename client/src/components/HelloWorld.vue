<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <button @click="Send">Send</button>
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
  props: {
    msg: String
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
