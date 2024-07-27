import {  useState } from "react";
import ShortScreen from "./ShortScreen";
import FormComponent from "../components/FormComponent";
import FormContainer from "../components/FormContainer";

const Home = () => {
    const [component, setComponent] = useState(true)
    const [url, setUrl] = useState("")

    const handleClick = (val) => {
        setComponent(false)
        setUrl(val);
    }

    const handleBack = () => {
        setComponent(true)
    }


    return(
        <FormContainer>
            <h1 style={{textAlign: "center"}}>Shorten your link</h1>
            { component ? 
                <FormComponent handleClick = {handleClick} />
                : <ShortScreen url={url} handleBack={handleBack}/> }
        </FormContainer>

    )
}

export default Home;