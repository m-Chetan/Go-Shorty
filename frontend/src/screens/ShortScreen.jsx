import { Button, Form, InputGroup } from "react-bootstrap";
import FormContainer from "../components/FormContainer";
import { useRef, useState } from "react";
import QRCodeGenerator from "../components/QRCodeGenerator"
import {Link} from "react-router-dom"

const ShortScreen = ({url,handleBack}) => {
    const [qr, setQr] = useState(false)
    const shortUrlInput = useRef(null)

    const handleClick = () => {
        shortUrlInput.current.select()
        navigator.clipboard.writeText(shortUrlInput.current.value)
            .then(() => alert(`copied: ${shortUrlInput.current.value}`))
            .catch((err) => console.log(err))

    }

    const handleQR = () => {
        setQr(true);
    }

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
                        <InputGroup>
                            <Form.Control
                                ref={shortUrlInput}
                                type="text" 
                                placeholder="Example: https://google.com" 
                                value="shorturl"
                                readOnly
                            />
                            <Button onClick={handleClick}>Copy</Button>
                        </InputGroup>
                    </Form.Group> 
                    <Form.Group>
                    <Button onClick={handleQR}>Generate QRCode</Button>
                    <Link href="/">
                        <Button className="mx-4" onClick={handleBack}>Shorten Another</Button>
                    </Link>
                    </Form.Group>
                    
                </Form>

                {qr && 
                    <div className="d-grid gap-auto">
                        <div className="mx-auto">
                            <QRCodeGenerator shortUrl={shortUrlInput.current.value} />
                        </div>
                        <Button className="my-3" >Download</Button>
                    </div>
                }

        </FormContainer>
    )
}


export default ShortScreen;
