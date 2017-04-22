const Message = function() {
  this.body = ''
}

Vue.component('message', {
  props: ['body'],
  template: '<div>{{ body }}</div>'
})

const app = new Vue({
  el: '#app',
  data: {
    messages: [],
    newMessage: new Message()
  },
  created() {
    this.getMessages();
  },
  methods: {
    getMessages() {
      fetch('/api/messages')
        .then(response => response.json())
        .then(data => {
          this.messages = data
        })
    },
    sendMessage() {
      const message = this.newMessage;
      fetch('/api/messages', {
        method: 'POST',
        body: JSON.stringify(message)
      })
      .then(response => response.json())
      .then(response => {
        if (response.error) {
          alert(response.error.message);
          return;
        }
        this.messages.push(response)
        this.clearMessage()
      }).catch(error => {
        console.log(error)
      })
    },
    clearMessage() {
      this.newMessage = new Message()
    }        
  }
})