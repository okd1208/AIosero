import axios from 'axios';export default {
  data () {
    return {
      info: '',
    }
  },
  methods: {
    async getPutDiscPosition() {
      // axios
      // .get('http://localhost:8888')
      // .then(response => (this.info = response))
      // console.log(this.info);
      axios.get('http://localhost:8888', {
        params: {
          userColor: 'white'
        }
      })
      .then(function (response) {
        console.log(response.data);
      });
      // await axios.request({
      //   method: 'get',
      //   url: 'http://localhost:8888',
      //   data: {
      //     firstName: 'Fred',
      //     lastName: 'Flintstone'
      //   }
      // });
    },
  },
}
