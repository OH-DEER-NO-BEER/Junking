using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Runtime.InteropServices;
using UnityEngine.UI;
using UnityEngine.SceneManagement;
using MiniJSON;

public class Enter_roomid : MonoBehaviour
{
    [SerializeField]
    private GameObject Get_Roomid;
    [SerializeField]
    private GameObject Error_text;

    private string RoomId;
    private string Responce_data;

    private string json_text = "";

    AudioSource audioSource;

    //JS -> SendMessage("EventControl","SetJson",string data)
    public void SetJson(string text)
    {
        json_text = text;
    }
    /*
    [DllImport("__Internal")]
    private static extern void GetLocalStorage(string key);
    */
    [DllImport("__Internal")]
    private static extern void SetLocalStorage(string key,string value);

    private void GameSceneLoaded(Scene next, LoadSceneMode mode)
    {
        var gameText = GameObject.Find("First_Canvas").transform.Find("Status").GetComponent<Text>();
        //var gameManager = GameObject.Find("EventControl").GetComponent<Recive_value>();
        //gameManager.player_status = Responce_data;
        gameText.text = "RoomID:" + RoomId;
        SceneManager.sceneLoaded -= GameSceneLoaded;
    }

    /*
    private string CheckInResponse_func(string data)
    {
        var json = Json.Deserialize(data) as Dictionary<string, object>;
        string CheckInResponse = (string)json["message"];

        return CheckInResponse;
    }
    */

    public void OnSendClick()
    {
        RoomId = Get_Roomid.GetComponent<Text>().text;
        audioSource.Play();

        if (RoomId == "")
        {
            Error_text.GetComponent<Text>().text = "Plese Input RoomID!!!";
            return;
        }
#if !UNITY_EDITOR
        SetLocalStorage("roomID", RoomId);
#endif

        SceneManager.LoadScene("Room_in_Scene"); //check
    }

    void Start()
    {
        audioSource = GameObject.Find("First_Canvas").transform.Find("Send").GetComponent<AudioSource>();
        SceneManager.sceneLoaded += GameSceneLoaded;
        //RoomId = Get_Roomid.GetComponent<Text>().text;
    }

    // Update is called once per frame
    void Update()
    {
        if (json_text == "")
        {
            return;
        }
        /*
        Responce_data = CheckInResponse_func(json_text);

        if(Responce_data == "Room Occupied")
        {
            Error_text.GetComponent<Text>().text = "An error has occurred";
            json_text = "";
        }
        else
        {
            SceneManager.LoadScene("Room_in_Scene");
        }
        */
    }
}
