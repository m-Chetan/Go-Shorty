import QRCode from "react-qr-code";

const QRCodeGenerator = ({shortUrl}) => {

    return(
        <QRCode value={shortUrl} size={200}/>
    )
}

export default QRCodeGenerator;