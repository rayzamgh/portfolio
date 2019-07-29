<template>
  <div class="font-lato bgimg" style="height: 110vh; overflow: hidden" scroll="no">
    <link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Raleway">
    <link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">      
    
    <header class="w3-display-container w3-wide bgimg w3-grayscale-min">
      <div class="w3-display-middle w3-text-white w3-center" style = "margin-top : 40vh">
        <h1 class="w3-jumbo">Rayza Mahendra</h1>
        <h2>Software Engineer in Training</h2>
        <h2><b>09.07.2019</b></h2>
      </div>
    </header>

    <!-- <img v-if="show" class="image" src = "https://i.kym-cdn.com/entries/icons/original/000/025/382/Screen_Shot_2018-02-06_at_3.37.14_PM.png"> -->
    <div class="w3-display-container bgimg2">
      <div class="w3-display-middle w3-text-white w3-center">
        <h1 class="w3-jumbo"></h1><br>
        <h2></h2>
      </div>
    </div>
    <div>
      <button @click="buttonPressedShowForm()" class="button" style="vertical-align:middle">
        <span>
          Input Address 
        </span>
      </button>
    </div>
    <form v-if="show">
      <div style="width: 200px;">
        PROVINSI
        <input style="padding: 3px 0 2px 10px; color:white; border-radius: 20px; background-color:transparent;" class = "formalamat" v-model = "provinsi">
      </div>
      <div style="width: 200px;">
        KOTA :
        <input style="padding: 3px 0 2px 10px; color:white; border-radius: 20px; background-color:transparent;" class = "formalamat" v-model = "kota">
      </div>
      <div style="width: 200px;">
        JALAN :
        <input style="padding: 3px 0 2px 10px; color:white; border-radius: 20px; background-color:transparent;" class = "formalamat" v-model = "jalan">
      </div>
      <div>
        <button @click="buttonPressed()" class="button" style="vertical-align:middle">
          <span>
          OK
          </span>
        </button>
      </div>
    </form>
    <div>
        <button class="buttonup" style="top:91%; left: 0%;">
           <div>
             About Me
             <img style="height:10vh; width:10vh;" src="https://memestatic.fjcdn.com/pictures/A+jotaro+kujo+character+study+jotaro+kujo+is+the+main_19a340_6325739.jpg">
               Currently Undergraduate in ITB
           </div>
        </button>
    </div>
    <div>
        <button class="buttonup" style="top:91%; left: 40%;">
          <div>
            Extra Pepega
            <iframe width="150" height="200"
            :src="linkyt" allow="autoplay; encrypted-media" allowfullscreen >
            </iframe>
          </div>
        </button>
    </div>
    <div>
      <button class="buttonupquery" style="top:91%; left: 55%;">
        <div>
          Query Youtube:
          <input
            style="padding: 3px 0 2px 10px; color:white; border-radius: 20px; background-color:transparent; "
            v-model = "queryYt"
            @keypress="checkEnter($event)">
          <button class="button" @click="buttonPressedYoutube()">
            Search
          </button>
        </div>
      </button>
    </div>
    <div>
      <button class="buttonup" style="top:91%; left:85%;">
        <div>
          Contact
        </div>
      </button>
    </div>
  <script src="https://code.jquery.com/jquery-2.1.3.min.js"></script>
  <script src="https://apis.google.com/js/client.js?onload=init"></script>
  </div> 
</template>
<script>
import axios from 'axios'
export default {
    data () {
        return {
            xx: '',
            show: false,
            jalan: '',
            kota: '',
            provinsi: '',
            queryYt: '',
            linkyt: ''
        }
    },
    watch: {
      modalOpen: function (isOpen) {
        if (isOpen) {
          document.documentElement.style.overflow = 'hidden'
          // document.documentElement is the same as using document.querySelector('#root')
        } else {
          document.documentElement.style.overflow = 'auto'
        }
      }
    },
    methods: {
      buttonPressed() {
        axios.post('http://localhost:8000/api/v1/alamat',{
          id: '',
          provinsi: this.provinsi,
          kota : this.kota,
          jalan : this.jalan
        })
        .then(function (response){
          alert("Data entry submitted")
          console.log(response);
        })
        .catch(function (response){
          alert("Server is inactive")
          console.log(response.message)
        })
      },

      buttonPressedShowForm() {
        this.show = !this.show
      },

      checkEnter($event) {
        if (event.keyCode === 13) {
          this.buttonPressedYoutubeManual()
        }
      },

      buttonPressedYoutube() {
        this.linkyt = 'https://www.youtube.com/embed?listType=search&list=' + this.queryYt
      },

      buttonPressedYoutubeManual() {
        let query = this.queryYt.split(' ').join('+')
        axios.get('http://localhost:8000/api/v1/utub?search=' + query).then(res => {
          if (res.data.message == "Valid") {
            this.linkyt = 'https://www.youtube.com/embed/'+ res.data.vidID  
            // this.linkyt = 'http://www.youtube.com/embed/videoseries?list=e98_o4Cun9Y-6iwugQvrp9M'
          }else{
            alert(res.data.message)
          }
        })

        .catch(function (error){
          console.log(error)
        })
      }
    }
}

</script>
<style>
#centered {
  text-align: center;
}
.input {
  color: blue;
  font-weight: 1000;
}
.pepego{
  color: red
}
.image {
  height: 30rem;
  width: 50rem;
}
.font-raleway {
  font-family: Raleway;
}
.font-lato {
  font-family: Lato;
  color: blanchedalmond;
}
body,h1,h2{font-family: "Raleway", sans-serif}
body, html {height: 100%}
p {line-height: 2}
.bgimg, .bgimg2 {
  min-height: cover;
  background-position: center;
  background-size: 100% 120%;
}
.bgimg {background-image: url("https://i.imgur.com/5bZbbC9.png")}
.bgimg2 {background-image: url("https://pm1.narvii.com/5910/22d0b5e1d282837e1d3b68bbd5beddb51bb1c730_hq.jpg")}

.button {
  display: inline-block;
  border-radius: 3px;
  background-color: transparent;
  border: none;
  color: #FFFFFF;
  text-align: center;
  font-size: 18px;
  padding: 20px;
  width: 200px;
  transition: all 0.5s;
  cursor: pointer;
  margin: 5px;
}

.button span {
  cursor: pointer;
  display: inline-block;
  position: relative;
  transition: 0.5s;
}

.button span:after {
  content: '\00bb';
  position: absolute;
  opacity: 0;
  top: 0;
  right: -20px;
  transition: 0.5s;
}

.button:hover span {
  padding-right: 25px;
}

.button:hover span:after {
  opacity: 1;
  right: 0;
}
html, body {
  overflow: hidden;
}
.buttonup {
  transition: 2s;
  position: absolute;
  padding: 0 0 100% 0 ;
  height: 200px;
  width: 150px;
  font-size: 17px;
  color: white;
  font-family: Raleway;
  background: transparent;
  border: none;
  transform-origin: top;
}

.buttonup div{
  transition: 2s;
}

.buttonup:hover{
  transform: translateY(-30vh)
}

.buttonupquery {
  transition: 2s;
  position: absolute;
  padding: 0 0 100% 0 ;
  height: 200px;
  width: 150px;
  font-size: 17px;
  color: white;
  font-family: Raleway;
  background: transparent;
  border: none;
  transform-origin: top;
}

.buttonupquery div{
  transition: 2s;
}
.buttonupquery:hover{
  transform: translateY(-10vh)
}
</style>