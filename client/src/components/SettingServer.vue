<template>
  <div id="app">
    <form @submit.prevent="submitForm">
      <div class="input-group">
        <label for="endpoint1">エンドポイント 1:</label>
        <input class="input-field" v-model="endpoints.endpoint1" id="endpoint1" type="text" placeholder="エンドポイント 1 を入力してください" required>
      </div>

      <div class="input-group">
        <label for="endpoint2">エンドポイント 2 (オプション):</label>
        <input class="input-field" v-model="endpoints.endpoint2" id="endpoint2" type="text" placeholder="エンドポイント 2 を入力してください">
      </div>

      <div class="input-group">
        <label for="interval">間隔 (ms):</label>
        <input class="input-field" v-model.number="interval" id="interval" type="number" min="0" placeholder="間隔を入力してください" required>
      </div>

      <button class="submit-btn" type="submit">保存</button>
    </form>
  </div>
</template>

<script>
import { mapMutations, mapActions } from 'vuex'

export default {
  computed: {
    endpoints: {
      get() {
        return this.$store.state.endpoints
      },
      set(endpoints) {
        this.$store.commit('setEndpoints', endpoints)
      },
    },
    interval: {
      get() {
        return this.$store.state.interval
      },
      set(interval) {
        this.$store.commit('setInterval', interval)
      },
    },
  },
  methods: {
    ...mapMutations(['setEndpoints', 'setInterval']),
    ...mapActions(['saveEndpoints', 'saveInterval']),
    submitForm() {
      this.saveEndpoints(this.endpoints);
      this.saveInterval(this.interval);
    },
  },
}
</script>

<style scoped>
form {
  width: clamp(400px, 60%, 2000px);
  margin: auto;
}

.input-group {
  margin-bottom: 20px;
}

.input-field {
  display: block;
  width: 100%;
  padding: 10px;
  font-size: 16px;
  border-radius: 5px;
  border: 1px solid #ccc;
}

.submit-btn {
  padding: 10px 20px;
  background-color: #0080ff;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.submit-btn:hover {
  background-color: #005cbf;
}
</style>
