import {  useState } from "react";
import ShortScreen from "./ShortScreen";
import FormComponent from "../components/FormComponent";
import FormContainer from "../components/FormContainer";

const Home = () => {
    const [component, setComponent] = useState(true)
    const [url, setUrl] = useState("")
    const [shortUrl, setShortUrl] = useState("")
    const [error, setError] = useState(null);

    const handleClick = (val) => {
        setComponent(false)
        setUrl(val);
        shortenUrl();
    }

    const shortenUrl = async() => {
        const request = {
            original_url: url
        };

        try {
            const response = await fetch('http://localhost:8080/shorten',{
                method: 'POST',
                headers: {
                  'Content-Type': 'application/json',
                },
                body: JSON.stringify(request),
            });
            if (!response.ok) {
              throw new Error('Network response was not ok');
            }
            const result = await response.json();
            setShortUrl(result.Short_Url);
        } catch (error) {
            setError(error.message);
        }
    }

    const handleBack = () => {
        setUrl("")
        setComponent(true)
    }


    return(
        <FormContainer>
            <h1 style={{textAlign: "center"}}>Shorten your link</h1>
            { component ? 
                <FormComponent handleClick = {handleClick} />
                : <ShortScreen url={url} shortUrl={shortUrl} handleBack={handleBack} error={error}/> }

        </FormContainer>

    )
}

export default Home;