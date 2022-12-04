<template>
  <div class="othelloBoard">
    <div v-for="row of 8" :key="row" class="row">
      <div v-for="col of 8" :key="col" class="col" @click="putDisc(row,col)">
        <div :id="getClassName(row,col)"></div>
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
      putPositions: {
        white: {
          1:[1],
          2:[],
          3:[6],
          4:[4],
          5:[5,8],
          6:[],
          7:[],
          8:[8],
        },
        black: {
          1:[],
          2:[],
          3:[],
          4:[5,6,7,8],
          5:[4],
          6:[],
          7:[],
          8:[],
        },
      },
      nextPutPositions: {
        white: {
          1:[],
          2:[],
          3:[],
          4:[],
          5:[],
          6:[],
          7:[],
          8:[],
        },
        black: {
          1:[],
          2:[],
          3:[],
          4:[],
          5:[],
          6:[],
          7:[],
          8:[],
        },
      }
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
    allSet() {
      const putPosi = this.putPositions;
      Object.keys(putPosi).forEach((color) => {
        const tmpRow = putPosi[color];
        Object.keys(tmpRow).forEach((row) => {
          tmpRow[row].forEach(col => {
            this.updateNextPosi(row, col, color);
            const element = document.getElementById(`row${row}-col${col}`);
            element.classList.add(color);
          })
        });
      });
    },
    updateNextPosi(row, col, color) {
      row = Number(row)
      console.log(`${color}:r${row}-c${col}`);
      const enemyColor = this.getReverseColor(color);
      for (let i = -1; i < 2 ; i++) {
        const r = row + i;
        if (r <= 0 || r > 8) {continue;}
        for (let n = -1; n < 2 ; n++) {
          const c = col + n;
          if (c <= 0 || c > 8){
            continue;
          }
          if (this.isExistDisc(r,c,enemyColor)) {
            if (row === r) {
              const lastCol = this.exploreHorizontal(r, c, enemyColor, c-col);
              if(lastCol) {
                this.nextPutPositions[color][r].push(lastCol);
              }
            } else if (col === c) {
              const lastRow = this.exploreVertical(r, c, enemyColor, r-row);
              if(lastRow) {
                this.nextPutPositions[color][lastRow].push(col);
              }
            } else {
              const lastPosi = this.exploreDiagonal(r, c, enemyColor, [r-row, c-col]);
              if(lastPosi["row"] && lastPosi["col"]) {
                this.nextPutPositions[color][lastPosi["row"]].push(lastPosi["col"]);
              }
            }
          }
        }
      }
    },
    isExistDisc(row,col,color) {
      const arr = this.putPositions[color][row];
      if(arr.indexOf(col) !== -1) {
        return true;
      }
      return false;
    },
    exploreHorizontal(row, col, color, i) {
      let lastCol = null;
      while(col > 1 && col < 8){
        col+=i;
        if (!this.isExistDisc(row, col, color)) {
          const rc = this.getReverseColor(color);
          if (!this.isExistDisc(row, col, rc)) {
            lastCol = col;
          }
          break;
        }
      }
      return lastCol;
    },
    exploreVertical(row, col, color, i) {
      let lastRow = null;
      while(row > 1 && row < 8){
        row+=i;
        if (!this.isExistDisc(row, col, color)) {
          const rc = this.getReverseColor(color);
          if (!this.isExistDisc(row, col, rc)) {
            lastRow = row;
          }
          break;
        }
      }
      return lastRow;
    },
    exploreDiagonal(row, col, color, i) {
      let lastRow = null;
      let lastCol = null;
      while((row > 1 && row < 8) && (col > 1 && col < 8)){
        row+=i[0];
        col+=i[1];
        if (!this.isExistDisc(row, col, color)) {
          const rc = this.getReverseColor(color);
          if (!this.isExistDisc(row, col, rc)) {
            lastRow = row;
            lastCol = col;
          }
          break;
        }
      }
      return {row: lastRow, col: lastCol};
    },
    getReverseColor (color) {
      return color === "white" ? "black" : "white";
    },
    putDisc(row,col) {
      console.log(row,col)
    }
  },
  mounted() {
    this.allSet();
  }
}
</script>

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
