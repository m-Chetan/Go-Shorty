import { Button, Form } from "react-bootstrap"
import FormContainer from "./FormContainer"
import {useState}  from "react";

const FormComponent = ({handleClick}) => {
    const [val,setVal] = useState("");

    const handleButtonClick = () =>{
        handleClick(val)
    }

    return(
        <FormContainer>
                <Form>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Label>Paste your link</Form.Label>
                        <Form.Control 
                            type="text" 
                            placeholder="Example: https://google.com" 
                            onChange = {(e) => setVal(e.target.value)}
                        />
                    </Form.Group>
                    <Button onClick = {handleButtonClick}>Shorten</Button>
                </Form>
        </FormContainer>
    )
}

export default FormComponent;