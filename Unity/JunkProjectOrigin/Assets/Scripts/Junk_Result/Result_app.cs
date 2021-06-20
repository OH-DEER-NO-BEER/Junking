using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class Result_app : MonoBehaviour
{
    [SerializeField]
    GameObject[] Hands_obj_type;

    GameObject myHand;
    GameObject enemyHand;

    [SerializeField]
    GameObject[] myhand_light;
    [SerializeField]
    GameObject[] enemyhand_light;

    [SerializeField]
    Text result_canvas;

    AudioSource audioSource;
    [SerializeField]
    AudioClip winclip;
    [SerializeField]
    AudioClip loseclip;

    [HideInInspector]
    public string self_h = "paper";
    [HideInInspector]
    public string oppenent_h = "rock";

    [HideInInspector]
    public string result_Junk = "win";

    public void OnConfirmClick()
    {
        SceneManager.LoadScene("Room_id_Scene");
    }

    // Start is called before the first frame update
    void Start()
    {
        myHand = Instantiate(Hands_obj_type[Hands_type.trans_number(self_h)]);
        myHand.transform.position = new Vector3(-5.5f, -3, 0);
        myHand.transform.eulerAngles = new Vector3(0, 90, 90);
        myHand.transform.localScale = new Vector3(25, 25, 25);

        enemyHand = Instantiate(Hands_obj_type[Hands_type.trans_number(oppenent_h)]);
        enemyHand.transform.position = new Vector3(5.5f, -3, 0);
        enemyHand.transform.eulerAngles = new Vector3(0, 90, 90);
        enemyHand.transform.localScale = new Vector3(25, 25, 25);

        audioSource = gameObject.GetComponent<AudioSource>();
        result_canvas.text = result_Junk;

        if (result_Junk == "win")
        {
            audioSource.clip = winclip;
            audioSource.Play();
        }
        else
        {
            audioSource.clip = loseclip;
            audioSource.Play();
        }

    }

    // Update is called once per frame
    void Update()
    {
        
    }
}
