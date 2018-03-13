<?php

class JsonRPC
{
    public $conn;

    function __construct($host, $port)
    {
        $this->conn = fsockopen($host, $port, $errno, $errstr, 3);
        if (!$this->conn) {
            return false;
        }
    }

    public function Call($method, $params)
    {
        $obj = new stdClass();
        $obj->code = 0;

        if (!$this->conn) {
            $obj->info = "jsonRPC-socket-tcp连接失败!";
            return $obj;
        }
        $err = fwrite($this->conn, json_encode(array(
                'method' => $method,
                'params' => array($params),
                'id' => 0,
            )) . "\n");
        if ($err === false) {
            fclose($this->conn);
            $obj->info = "jsonRPC发送参数失败!socket-tcp资源是否释放";
            return $obj;
        }

        stream_set_timeout($this->conn, 0, 3000);
        $line = fgets($this->conn);
        fclose($this->conn);
        if ($line === false) {
            $obj->info = "jsonRPC返回消息为空!请检查自己的rpc-client代码";
            return $obj;
        }
        $temp = json_decode($line);
        $obj->code = $temp->error == null ? 1 : 0;
        $obj->data = $temp->result;
        return $obj;
    }
}


function json_rpc_ip_address($ipString)
{
    $client = new JsonRPC("127.0.0.1", 3344);
    $obj = $client->Call("Ip2addr.Address", ['IpString' => $ipString]);
    return $obj;
}