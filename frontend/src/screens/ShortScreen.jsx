import { Button, Form, InputGroup } from "react-bootstrap";
import FormContainer from "../components/FormContainer";
import { useRef } from "react";
// import QRCodeGenerator from "../components/QRCodeGenerator"
import {Link, useNavigate} from "react-router-dom"

const ShortScreen = ({url,shortUrl,handleBack, error}) => {
    //const [qr, setQr] = useState(false)
    const shortUrlInput = useRef(null)
    const navigate = useNavigate()

    const handleClick = () => {
        shortUrlInput.current.select()
        navigator.clipboard.writeText(shortUrlInput.current.value)
            .then(() => alert(`copied: ${shortUrlInput.current.value}`))
            .catch((err) => console.log(err))

    }

    const handleRedirect = () => {
        navigate(shortUrl)
    }

    // const handleQR = () => {
    //     setQr(true);
    // }

    return(
        <FormContainer >
                <Form className="my-3">
                    <Form.Group className="mb-3">
                        <Form.Label>Your Long URL</Form.Label>
                        <Form.Control 
                            type="text" 
                            value={url}
                            placeholder="Example: https://google.com" 
                            readOnly
                        />
                    </Form.Group>

                    <Form.Group className="mb-3">
                        <Form.Label>ShortURL</Form.Label>
                        {error && <p style={{color: "red"}}>Error: {error}</p>}
                        <InputGroup>
                            <Form.Control
                                ref={shortUrlInput}
                                type="text" 
                                value={shortUrl}
                                readOnly
                            />
                            <Button onClick={handleClick}>Copy</Button>
                        </InputGroup>
                    </Form.Group> 
                    <Form.Group>
                    {/* <Button onClick={handleQR}>Generate QRCode</Button> */}
                    <Link href="/">
                        <Button  onClick={handleBack}>Shorten Another</Button>
                        <Button className="mx-4" onClick={handleRedirect}>Redirect</Button>
                    </Link>
                    </Form.Group>
                    
                </Form>

                {/* {qr && 
                    <div className="d-grid gap-auto">
                        <div className="mx-auto">
                            <QRCodeGenerator shortUrl={shortUrlInput.current.value} />
                        </div>
                        <Button className="my-3" >Download</Button>
                    </div>
                } */}

        </FormContainer>
    )
}


export default ShortScreen;
