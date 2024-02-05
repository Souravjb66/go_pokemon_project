var call=async()=>{
    try{
        const url=await fetch("http://127.0.0.1:3000/getId/1")
        const response=await url.json()
        console.log("clicked s")
        console.log(response)
        const val=JSON.stringify(response)
        console.log(val)
        console.log(response.name)
        console.log(response.fetch.name)
        

    }catch(error){
        console.log("geting "+error)
    }
    


}
call()


//let arr='{"name":"sourav","id":22}'
        //const mol=JSON.parse(arr)