new Vue({
    el: '#Khufiya',
    data: {
        websocket: null, // Websocket Details
        msg: '', // Messages to be sent
        chatHistory: '', // Chat History
        email: null, // Identifying Email-ID for future
        username: null, // Username
        joined: false // Registration Details
    },

    created: function() {
        var self = this;
        this.websocket = new WebSocket('ws://' + window.location.host + '/websocket');
        this.websocket.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            self.chatHistory = self.chatHistory+'<div class="chip">'
                    + '<img src="' + self.gravatarURL(msg.email) + '">'
                    + msg.username
                + '</div>'
                + emojione.toImage(msg.message) + '<br/>';

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight;
        });
    },

    methods: {
        send: function () {
            if (this.msg!= ''){
                this.websocket.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: $('<p>').html(this.msg).text()
                    }
                ));
                this.msg = '';
            }
        },

        join: function () {
            if (!this.email) {
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.email = $('<p>').html(this.email).text();
            this.username = $('<p>').html(this.username).text();
            this.joined = true;
        },

        gravatarURL: function(email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        }
    }
});
