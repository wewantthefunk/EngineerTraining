var app = new Vue({
  el: '#app',
  data: {
    message: 'Welcome Engineering Training',
    result: '',
    responseAvailabe: false
  },
  methods: {
   fetchAPIData() {
      this.responseAvailable = false;
      fetch("/db", {
          "method": "GET",
          "headers": {
          }
      })
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
