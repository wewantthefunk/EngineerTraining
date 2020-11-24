var app = new Vue({
  el: '#app',
  data: {
    message: 'Welcome Engineering Training',
    result: '',
    responseAvailable: false,
    apiUrl: "/db",
    apiObject: {"method":"GET","headers":{}}
  },
  methods: {
   init(url, obj) {
      this.apiUrl = url;
      this.apiObject = obj;
   },
   fetchAPIData() {
      this.responseAvailable = false;
      fetch(this.apiUrl, this.apiObject)
      .then(response => {
          if(response.ok){
              return response.json()
          } else{
              alert("Server returned " + response.status + " : " + response.statusText);
          }
      })
      .then(response => {
          this.result = response.users;
          this.responseAvailable = true;
      })
      .catch(err => {
          console.log(err);
      })
   } //fetchAPIData
  } //methods
});
