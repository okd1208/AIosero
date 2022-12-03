<template>
  <div class="othelloBoard">
    <div v-for="n of 8" :key="n" class="row">
      <div v-for="i of 8" :key="i" class="col">
        <div :id="getClassName(n,i)"></div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'HelloWorld',
  data: function () {
    return {
      info: null,
      white: {
        4:[4],
        5:[5]
      },
      black: {
        4:[5],
        5:[4]
      },
    }
  },
  methods: {
    getClassName(row, col) {
      return `row${row}-col${col}`
    },
    async Send() {
      axios
      .get('http://localhost:8888')
      .then(response => (this.info = response))
      console.log(this.info);
    },
    allSet(arr, color) {
      Object.keys(arr).forEach(function (row) {
        arr[row].forEach(col => {
          console.log(`row${row}-col${col}`);
          var element = document.getElementById(`row${row}-col${col}`);
          element.classList.add(color);
        })
    });
    },
  },
  mounted() {
    this.allSet(this.white, "white");
    this.allSet(this.black, "black");
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.othelloBoard {
  color: aliceblue;
  margin: auto;
  width: 500px;
  height: 500px;
  background-color: rgb(25, 102, 76);
  border: 8px solid black;
}
.row {
  height: 12.5%;
  display: flex;
}
.col {
  width: 12.5%;
  text-align: center;
  border: 1px solid black;
  /* 子要素でmargin-autoが効くようにするため */
  display: flex;
}
.white,.black {
  margin: auto;
  height: 90%;
  width: 90%;
  border-radius: 50%;
}
.white {
  background-color: white;
}
.black {
  background-color: black;
}
</style>
