<template>
<div>
<br><br><br>
    <div class="cardregister">
        <h1>register</h1>
        <form action="" class="rigister__form" @submit="submitForm">
            <input type="text" class="rigister__input" placeholder="Username" v-model="username"><br>
            <input type="email" class="rigister__input" placeholder="email" v-model="gamil"><br>
            <input type="password" class="rigister__input" placeholder="password" v-model="password"><br>
            <input type="password" class="rigister__input" placeholder="repassword" v-model="repassword"><br>
            <button type="submit" class="rigister__button" :disabled="rigistering">submit</button>
        </form>
        <nuxt-link to="/login" class="aaa">login</nuxt-link>
    </div>
    </div>
</template>
<script lang="ts">
const FORMSPARK_ACTION_URL:string  = "https://submit-form.com/your-form-id";

export default {
    head(){
      return{
  title:["AX47-register"],
  link: [

    { rel: 'stylesheet', href: 'static/css/pages/register.css' }
  ]
} 
    },
    methods: {
    async submitForm() {
    //   await fetch(FORMSPARK_ACTION_URL, {
    //     method: "POST",
    //     headers: {
    //       "Content-Type": "application/json",
    //       Accept: "application/json",
    //     },
    //     body: JSON.stringify({
    //       message: this.message,
    //     }),
    //   });
      alert("Form"+this.username+"\n"+this.gamil+"\n"+this.password+"\n"+this.repassword+"\n");
    },
  },
}
</script>
