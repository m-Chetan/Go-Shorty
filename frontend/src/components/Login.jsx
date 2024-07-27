import { Button, Col, Form, Row } from "react-bootstrap";

const Login = ()=> {
    return (
        <Form>
            <Row className="justify-content-md-center mt-3">
                <Col xs={12} md={6}>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                    <Form.Label>Email address</Form.Label>
                    <Form.Control type="email" placeholder="Enter email" />
                    </Form.Group>
                
                    <Form.Group className="mb-3" controlId="formBasicPassword">
                    <Form.Label>Password</Form.Label>
                    <Form.Control type="password" placeholder="Password" />
                    </Form.Group>
                
                    <Button variant="primary" type="submit">
                        Submit
                    </Button>
                </Col>
            </Row>    
        </Form>
    )
}

export default Login;