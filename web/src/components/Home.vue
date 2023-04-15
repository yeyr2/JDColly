<template>
  <div>
    <div class="top">
      <div class="round">
        <div class="container">
          <img src="../assets/img/iPhone.jpg" />
          <img src="../assets/img/iPad2.jpg" />
          <img src="../assets/img/watch.jpg" />
          <img src="../assets/img/AirPods.jpg" />
          <img src="../assets/img/iPhone.jpg" />
          <img src="../assets/img/iPhone.jpg" />
        </div>
      </div>
    </div>
    <div class="main">
      <div class="iPhone box">
        <div class="interLink">
          <div class="name">iPhone 14 Pro</div>
        </div>
      </div>
      <div class="watch box">
        <div class="interLink">
          <div class="name">watch</div>
        </div>
      </div>
      <div class="iPad box">
        <div class="interLink">
          <div class="name">iPad Pro</div>
        </div>
      </div>
      <div class="AirPods box">
        <div class="interLink">
          <div class="name">AirPods</div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {};
</script>
<style scoped>
.top {
  width: 100%;
  height: 200px;
  background-color: #000000;
}
.round {
  perspective: 6000px;
  position: absolute;
  top: 42%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.container {
  position: absolute;
  top: 70%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 270px;
  height: 300px;
  transform-style: preserve-3d;
  animation-name: rotating;
  animation-duration: 15s;
  animation-timing-function: linear;
  animation-iteration-count: infinite;
  animation-play-state: running;
}

.container img {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  transition: all 0.3s;
}

.container img:hover {
  box-shadow: 0 0 8px 1px #24b4f3;
}

.container img:nth-of-type(1) {
  transform: rotateY(60deg) translateX(300px) rotateY(90deg);
}

.container img:nth-of-type(2) {
  transform: rotateY(120deg) translateX(300px) rotateY(90deg);
}

.container img:nth-of-type(3) {
  transform: rotateY(180deg) translateX(300px) rotateY(90deg);
}

.container img:nth-of-type(4) {
  transform: rotateY(2400deg) translateX(300px) rotateY(90deg);
}

.container img:nth-of-type(5) {
  transform: rotateY(300deg) translateX(300px) rotateY(90deg);
}

.container img:nth-of-type(6) {
  transform: rotateY(360deg) translateX(300px) rotateY(90deg);
}
@keyframes rotating {
  0% {
    transform: translate(-50%, -50%) rotateY(0deg);
  }

  100% {
    transform: translate(-50%, -50%) rotateY(360deg);
  }
}
.main {
  width: 100%;
  overflow: hidden;
}
.main .box {
  width: 100%;
  height: 580px;
}
.main .iPhone {
  background: url(../assets/img/iPhone.jpg) no-repeat center bottom #000000;
  margin-bottom: 1px;
}
.main .iPhone .interLink {
  position: relative;
  width: 800px;
  height: 100px;
  top: 10%;
  left: 50%;
  margin-left: -400px;
  text-align: center;
}
.main .iPhone .interLink .name {
  color: #ffffff;
  font-size: 48px;
}
.main .iPhone .interLink .turn {
  display: inline-block;
  color: blue;
  font-size: 21px;
}
.main .iPad {
  background: url(../assets/img/iPad.jpg) no-repeat center center;
  height: 580px;
  margin-bottom: 1px;
}
.main .iPad .interLink {
  position: relative;
  width: 400px;
  height: 200px;
  left: 10%;
  top: 40%;
  margin-left: -100px;
  text-align: center;
}
.main .iPad .interLink .name {
  color: #ffffff;
  font-size: 56px;
}
.main .iPad .interLink .turn {
  display: inline-block;
  color: blue;
  font-size: 21px;
}
.main .watch {
  background: url(../assets/img/watch.jpg) no-repeat center center;
}
.main .watch .interLink {
  position: relative;
  width: 400px;
  height: 200px;
  left: 10%;
  top: 40%;
  margin-left: -100px;
  text-align: center;
  vertical-align: middle;
}
.main .watch .interLink .name {
  color: #1d1d1f;
  font-size: 56px;
}
.main .watch .interLink .turn {
  display: inline-block;
  color: blue;
  font-size: 21px;
}
.main .AirPods {
  background: url(../assets/img/AirPods.jpg) no-repeat center center;
  /* height: 100%; */
}
.main .AirPods .interLink {
  position: relative;
  width: 400px;
  height: 200px;
  left: 46%;
  top: 32%;
  margin-left: -200px;
  text-align: center;
}
.main .AirPods .interLink .name {
  color: #1d1d1f;
  font-size: 120px;
}
.main .AirPods .interLink .turn {
  display: inline-block;
  color: blue;
  font-size: 21px;
}
</style>