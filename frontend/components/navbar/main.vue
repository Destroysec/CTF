<template>

<nav>
    <h1 class="header">
        DestroySEC
    </h1>
<!--     
    <ul class="main">
      
        
      
      <li>
            <NuxtLink to="/" >Home</NuxtLink>

      </li>
      <li>
        
            <NuxtLink to="/server">server</NuxtLink>

      </li>
<li>
        
            <NuxtLink to="/server">about</NuxtLink>

      </li>
      </ul> -->
          
        <ul class="aun">
          <li>
            <NuxtLink to="/login" class="login">
            <button>
login
            </button>
            
            
          </NuxtLink>
          </li>
          <li>
            <NuxtLink to="/register" class="register">
            <button>
register
            </button>
            
            
          </NuxtLink>
          </li>
        </ul>
          
 
</nav>



</template>
<script lang="ts">
export default ({
   head () {
return {
  link: [

    { rel: 'stylesheet', href: 'static/css/components/navbar/index.css' }
  ]
} 
}}
)
</script>
